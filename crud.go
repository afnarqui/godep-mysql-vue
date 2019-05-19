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
	//dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	//db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/defaultdb?sslmode=disable")
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/test?sslmode=disable")
	//db, err := sql.Open("postgres", dsn)

	    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    // Create the "accounts" table.
    if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
        log.Fatal(err)
    }

    // Insert two rows into the "accounts" table.
    if _, err := db.Exec(
        "INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"); err != nil {
        log.Fatal(err)
    }

    // Print out the balances.
    rows, err := db.Query("SELECT id, balance FROM accounts")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("Initial balances:")
    for rows.Next() {
        var id, balance int
        if err := rows.Scan(&id, &balance); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d %d\n", id, balance)
    }

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