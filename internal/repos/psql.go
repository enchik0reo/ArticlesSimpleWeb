package repos

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(c Config) (*sql.DB, error) {
	cnf := fmt.Sprintf("user=%s password=%s port=%s  dbname=%s sslmode=%s",
		c.Username, c.Password, c.Port, c.DBName, c.SSLMode)
	db, err := sql.Open("postgres", cnf)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	script := readScript()

	_, err = db.Exec(script)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func readScript() string {
	var filename = "docs/script.txt"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(b)
}
