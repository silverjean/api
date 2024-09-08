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

func (repo Post) FindPostByID(postID uint64) (models.Post, error) {
	lines, err := repo.db.Query(
		`SELECT P.*, U.nick FROM
		posts P INNER JOIN users U
		ON U.id = P.author_id WHERE P.id = ?`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer lines.Close()

	var post models.Post

	if lines.Next() {
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreateAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}