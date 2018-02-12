package handlers

import (
	"database/sql"
	"github.com/acaporrini/rest-api/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetPersons(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPersons(db))
	}
}

func GetPerson(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		return c.JSON(http.StatusOK, models.GetPerson(db, id))
	}
}

func CreatePerson(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.CreatePerson(db, c))
	}
}
