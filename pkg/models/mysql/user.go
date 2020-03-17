package mysql

import (
	"database/sql"
	"github.com/ihaveint/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil

}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil

}
