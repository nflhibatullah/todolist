package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"todolist/configs"
	"todolist/delivery/controllers/user"
	"todolist/delivery/routes"
	_userRepo "todolist/repository/user"
	"todolist/utils"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := _userRepo.New(db)
	UserController := user.New(userRepo)

	e := echo.New()
	routes.RegisterPath(e, UserController)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
