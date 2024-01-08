package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	articles, err := h.repo.GetAll()
	if err != nil {
		fmt.Fprintf(w, "Ошибка получения статей")
	}

	t.ExecuteTemplate(w, "index", articles)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "create", nil)
}

func (h *Handler) saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")

	if title == "" || anons == "" || fullText == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {
		err := h.repo.Save(title, anons, fullText)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Ошибка сохранения статьи")
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func (h *Handler) show(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	userId := vars["id"]
	h.id, err = strconv.Atoi(userId)
	if err != nil {
		fmt.Fprintf(w, "Ошибка Id")
	} else {
		showPost, err := h.repo.GetById(h.id)
		if err != nil {
			fmt.Fprintf(w, "Ошибка получения статьи")
		}

		t.ExecuteTemplate(w, "show", showPost)
	}
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	err := h.repo.DeleteById(h.id)
	if err != nil {
		fmt.Fprintf(w, "Ошибка удаления статьи")
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
