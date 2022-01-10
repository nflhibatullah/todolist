package user

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
