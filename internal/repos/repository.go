package repos

import "database/sql"

type Articles interface {
	GetAll() ([]Article, error)
	Save(title, anons, fullText string) error
	GetById(userId int) (Article, error)
	DeleteById(userId int) error
}

type Repository struct {
	Articles
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Articles: NewArticlesPostgres(db),
	}
}
