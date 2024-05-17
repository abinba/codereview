package repo

import (
	"github.com/abinba/codereview/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(username, password, first_name, last_name string) error {
	user := model.User{
		Username: username,
		Password: password,
		FirstName: first_name,
		LastName: last_name,
	}
	return repo.db.Create(&user).Error
}
