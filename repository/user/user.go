package user

import (
	"gorm.io/gorm"
	"todolist/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {

	if err := ur.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil

}

func (ur *UserRepository) Login(email, password string) {

}
