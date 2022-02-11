package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/shefchenkornd/rest_api/internal/models"
	"log"
)

// UserRepository Хотим, чтобы наше приложение общалось с моделью User через репозиторий UserRepository
type UserRepository struct {
	storage *Storage
}

var (
	tableUser = "users"
)

// Create user
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, &u.Password).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByLogin user by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	founded := false

	query := fmt.Sprintf("SELECT id, login FROM %s WHERE login=$1", tableUser)
	var userId int
	var userLogin string
	if err := ur.storage.db.QueryRow(query, login).Scan(&userId, &userLogin); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}

		return nil, false, err
	}
	founded = true

	user := models.User{
		Id:    userId,
		Login: userLogin,
	}

	return &user, founded, nil
}

// SelectAll all users
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 100)
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(user.Id, user.Login, user.Password); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
