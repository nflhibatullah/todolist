package user

import "todolist/entities"

type User interface {
	Register(entities.User) (entities.User, error)
}
