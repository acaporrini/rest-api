package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Person struct {
	Id         int
	First_Name string
	Last_Name  string
}

type PersonCollection struct {
	Persons []Person `json:"items"` // TODO Look better for this
}

func GetPersons(db *sql.DB) PersonCollection {
	var (
		person            Person
		personsCollection PersonCollection
	)

	sql := "SELECT * FROM persons"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
		personsCollection.Persons = append(personsCollection.Persons, person)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	defer rows.Close()

	return personsCollection
}

func GetPerson(db *sql.DB, id int) Person {
	var person Person
	sql := "SELECT * FROM persons where id = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)

	if err != nil {
		panic(err)
	}

	return person

}

func CreatePerson(db *sql.DB, c echo.Context) Person {
	var person Person
	sql := "INSERT INTO persons (first_name, last_name) VALUES (?,?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(c.FormValue("first_name"), c.FormValue("last_name"))

	if err2 != nil {
		panic(err)
	}

	row := db.QueryRow("select * from persons order by id desc limit 1")
	err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)

	return person
}
