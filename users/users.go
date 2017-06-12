package users

import (
	"database/sql"
	"edms-go-echo/models"
	"net/http"

	"github.com/labstack/echo"
)

// Handler ...
type Handler map[string]interface{}

// CreateUser ...
func CreateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		c.Bind(&user)
		id, err := models.CreateUser(db, user)
		if err == nil {
			return c.JSON(http.StatusCreated, Handler{
				"created": id,
			})
		}
		return err
	}
}
