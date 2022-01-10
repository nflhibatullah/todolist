package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Todo []TodoList
}
