package main

import (
	// "bytes"
	"database/sql"
	"fmt"
	// "net/http"

	"github.com/acaporrini/rest-api/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	e := echo.New()
	e.GET("/persons", handlers.GetPersons(db))
	e.GET("/persons/:id", handlers.GetPerson(db))
	e.POST("/persons", handlers.CreatePerson(db))

	e.Start(":8080")
}
