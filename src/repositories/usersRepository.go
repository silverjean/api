package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
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

func (user usersRepository) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick= fmt.Sprintf("%%%s%%", nameOrNick) // -> %nameOrNick% 

	lines, err := user.db.Query(
		"SELECT id, name, nick, email, create_date FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (user usersRepository) FindOne(ID uint64) (models.User, error) {
	lines, err := user.db.Query(
		"SELECT id, name, nick, email, create_date FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return 	models.User{}, err
	}
	defer lines.Close()
	
	var userModel models.User
	if lines.Next() {
		if err = lines.Scan(
			&userModel.ID,
			&userModel.Name,
			&userModel.Nick,
			&userModel.Email,
			&userModel.CreateDate,
		); err != nil {
			return models.User{}, errors.New("user not found")
		}
	}

	return userModel, nil
}

func (user usersRepository) UpdateUser(ID uint64, userBody models.User) error {
	statemant, err := user.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statemant.Close()

	if _, err = statemant.Exec(userBody.Name, userBody.Nick, userBody.Email, ID); err != nil {
		return err
	}

	return nil

}

func (user usersRepository) Delete(ID uint64) error {
	statement, err := user.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}