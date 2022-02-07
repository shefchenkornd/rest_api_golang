package storage

import (
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

// Find user by login
func (ur *UserRepository) Find(login string) (*models.User, bool, error) {
	var user *models.User
	founded := false

	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1", tableUser)
	if err := ur.storage.db.QueryRow(query, login).Scan(&user); err != nil {
		return nil, false, err
	}
	founded = true

	return user, founded, nil
}

// Select all users
func (ur *UserRepository) Select() ([]*models.User, error) {
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
