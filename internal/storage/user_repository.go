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

	query := fmt.Sprintf("SELECT id, login, password FROM %s WHERE login=$1", tableUser)
	var userId int
	var userLogin string
	var userPassword string
	if err := ur.storage.db.QueryRow(query, login).Scan(&userId, &userLogin, &userPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}

		return nil, false, err
	}
	founded = true

	user := models.User{
		Id:    userId,
		Login: userLogin,
		Password: userPassword,
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

func (ur UserRepository) Auth(u *models.User) (bool, error) {
	found := false
	var user models.User

	query := fmt.Sprintf("SELECT login FROM %v where login=$1 and password=$2", tableUser)
	err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&user.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return found, nil
		}

		return found, err
	}
	found = true

	return found, nil
}