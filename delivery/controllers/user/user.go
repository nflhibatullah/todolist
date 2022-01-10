package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"todolist/configs"
	"todolist/delivery/middlewares"
	"todolist/entities"
	"todolist/repository/user"
)

type UserController struct {
	Repo user.User
}

func New(user user.User) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := entities.User{}
		c.Bind(&tmp)
		hash, _ := bcrypt.GenerateFromPassword([]byte(tmp.Password), 14)

		UserData := entities.User{
			Name:     tmp.Name,
			Email:    tmp.Email,
			Password: string(hash),
		}

		res, err := uc.Repo.Register(UserData)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "Ada kesalahan dalam input")
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "Berhasil create user",
				"data":    res,
			},
		)
	}

}

func (uc UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpUser := entities.User{}

		if err := c.Bind(&tmpUser); err != nil {
			return err
		}

		res, err := uc.Repo.Login(tmpUser.Email, tmpUser.Password)
		if err != nil {
			log.Error(err)
			return c.JSON(
				http.StatusNotFound, map[string]string{
					"message": "Ada Kesalahan dalam login",
				},
			)
		}

		res.Token, _ = middlewares.CreateToken(int(tmpUser.ID), configs.SecretKey)

		if res.ID == 0 && res.Email == "" {
			return c.JSON(
				http.StatusNotFound, map[string]string{
					"message": "Ada Kesalahan dalam login",
				},
			)
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "Berhasil login",
				"data":    res,
			},
		)
	}

}
