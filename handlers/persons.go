package handlers

import (
	"database/sql"
	"github.com/acaporrini/rest-api/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
)

func GetPersons(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPersons(db))
	}
}
