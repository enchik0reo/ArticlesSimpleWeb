package repos

import "database/sql"

type ArticlesPostgres struct {
	db *sql.DB
}

type Article struct {
	Id                     int
	Title, Anons, FullText string
}

func NewArticlesPostgres(db *sql.DB) *ArticlesPostgres {
	return &ArticlesPostgres{
		db: db,
	}
}

func (r *ArticlesPostgres) GetAll() ([]Article, error) {
	row, err := r.db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	articles := []Article{}
	for row.Next() {
		a := Article{}
		err = row.Scan(&a.Id, &a.Title, &a.Anons, &a.FullText)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func (r *ArticlesPostgres) Save(title, anons, fullText string) error {
	_, err := r.db.Exec("INSERT INTO articles (title, anons, full_text) VALUES ($1, $2, $3)", title, anons, fullText)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticlesPostgres) GetById(userId int) (Article, error) {
	row, err := r.db.Query("SELECT * FROM articles WHERE id = $1", userId)
	if err != nil {
		return Article{}, err
	}
	defer row.Close()

	showPost := Article{}

	for row.Next() {
		a := Article{}
		err = row.Scan(&a.Id, &a.Title, &a.Anons, &a.FullText)
		if err != nil {
			return Article{}, err
		}
		showPost = a
	}

	return showPost, nil
}

func (r *ArticlesPostgres) DeleteById(userId int) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id = $1", userId)
	return err
}
