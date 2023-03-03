package handler

import (
	"ArticlesLittleWeb/internal/repos"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	repo *repos.Repository
	id   int
}

func New(repository *repos.Repository) *Handler {
	return &Handler{
		repo: repository,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", h.index).Methods("GET")
	rtr.HandleFunc("/create", h.create).Methods("GET")
	rtr.HandleFunc("/save_article", h.saveArticle).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", h.show).Methods("GET")
	rtr.HandleFunc("/delete_article", h.delete).Methods("POST")

	return rtr
}
