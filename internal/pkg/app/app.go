package app

import (
	"ArticlesLittleWeb/internal/handler"
	"ArticlesLittleWeb/internal/repos"
	"ArticlesLittleWeb/internal/server"
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *server.Server
	db         *sql.DB
	repo       *repos.Repository
	handler    *handler.Handler
}

func New() *App {
	a := &App{}

	log.SetFormatter(new(log.JSONFormatter))

	if err := initConfig(a); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	a.httpServer = server.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repos.NewPostgresDB(repos.Config{
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to inicialise DB: %s", err.Error())
	}

	a.db = db

	a.repo = repos.NewRepository(a.db)

	a.handler = handler.New(a.repo)

	log.Println("New application instance created")
	return a
}

func (a *App) Run() {
	go func() {
		if err := a.httpServer.Run(viper.GetString("port"), a.handler.InitRoutes()); err != nil {
			log.Fatalf("error occurred while running http server: %v", err)
		}
	}()

	log.Print("ArticlesApp Successfully Started ")

	go func() {
		for {
			for _, s := range []string{".   ", "..  ", "... ", "....", " ...", "  ..", "   ."} {
				fmt.Printf("\r%s", s)
				time.Sleep(time.Millisecond * 150)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("ArticlesApp Shutting Down ")

	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		log.Error("error occured on server shutting down: %s", err.Error())
	}

	if err := a.db.Close(); err != nil {
		log.Error("error occured on db connection close: %s", err.Error())
	}
}

func initConfig(a *App) error {
	viper.AddConfigPath("config")
	viper.SetConfigName("cnf")
	return viper.ReadInConfig()
}
