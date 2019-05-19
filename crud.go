package main

import (
	"log"
	"fmt"
	"time"
	"errors"
)
// struct student

type Estudiante struct {
	ID int
	Name string
	Age int16
	Active bool
	CreatedAt time.time
	UpdatedAd time.Time
}

// create or add student

func create (e Estudiante) error {
	q := `INSERT INTO 
				estudiantes(Name, Age, Active)
				VALUES ($1,$2,$3)
			`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active)

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Should error rows")
	}

	return nil
}

func main() {
	e := Estudiante {
		Name: "Alejandro",
		Age: 30,
		Active: true
	}

	err := Create(e)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create with exited")
}