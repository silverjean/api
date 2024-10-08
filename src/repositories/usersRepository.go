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

func (user usersRepository) FindByEmail(email string) (models.User, error) {
	line, err := user.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var userModel models.User

	if line.Next() {
		if err = line.Scan(&userModel.ID, &userModel.Password); err != nil {
			return models.User{}, err
		}


	}
	return userModel, nil
}

func (user usersRepository) Follow(userID, followerID uint64) error {
	statement, err := user.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (user usersRepository) Unfollow(userID, followerID uint64) error {
	statement, err := user.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (user usersRepository) FindFollowers(userID uint64) ([]models.User, error) {
	lines, err := user.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.create_date 
			FROM users u 
			INNER JOIN followers f ON u.id = f.follower_id 
				WHERE f.user_id = ? `,
		userID,
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

func (user usersRepository) FindFollowing(ID uint64) ([]models.User, error) {
	lines, err := user.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.create_date 
			FROM users u 
			INNER JOIN followers f ON u.id = f.user_id 
				WHERE f.follower_id = ? `,
		ID,
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

func (repo usersRepository) FindPassByID(userID uint64) (string, error) {
	line, err := repo.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repo usersRepository) UpdatePass(userID uint64, pass string) error {
	statemant, err := repo.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err = statemant.Exec(pass, userID); err != nil {
		return err
	}

	return nil
}