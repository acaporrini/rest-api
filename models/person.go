package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
