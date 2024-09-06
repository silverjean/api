package repositories

import (
	"api/src/models"
	"database/sql"
)

type Post struct {
	db *sql.DB
}

func NewRepositoryPost(db *sql.DB) *Post {
	return &Post{db}
}

func (repo Post) Create(post models.Post) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInserID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInserID), nil
}