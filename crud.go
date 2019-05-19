package main

import (
	"log"
	"fmt"
	"time"
	"errors"
	"database/sql"
	_ "github.com/lib/pq"
)

//connection should conect with bd
func getConnection() *sql.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	
	return db
}

// struct student

type Estudiante struct {
	ID int
	name string
	age int16
	active bool
	CreatedAt time.Time
	UpdatedAd time.Time
}

// create or add student

func create (e Estudiante) error {
	q := `INSERT INTO 
				estudiantes(Name, Age, Active)
				VALUES ($1,$2,$3)`
			

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.name, e.age, e.active)

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
		name: "Alejandro",
		age: 30,
		active: true,
	}

	err := create(e)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create with exited")
}