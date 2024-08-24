package repositories

import (
	"api/src/models"
	"database/sql"
)

type usersRepository struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *usersRepository {
	return &usersRepository{db}
}

func (userRepo usersRepository) Create(userModel models.User) (uint64, error) {
	statement, err := userRepo.db.Prepare(
		"insert into users (name, nick, email, password) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(userModel.Name, userModel.Nick, userModel.Email, userModel.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}