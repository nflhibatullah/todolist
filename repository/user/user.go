package user

import (
	"golang.org/x/crypto/bcrypt"
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

func (ur *UserRepository) Login(email, password string) (entities.User, error) {

	var err error
	foundUser := entities.User{}

	getPass := entities.User{}
	ur.db.Select("password").Where("Email = ?", email).Find(&getPass)
	err = bcrypt.CompareHashAndPassword([]byte(getPass.Password), []byte(password))
	if err != nil {
		return entities.User{}, err
	}

	ur.db.Select("ID", "name", "email").Where("Email = ?", email).Find(&foundUser)
	return foundUser, nil
}
