package main

import (
	"./estudiante.go"
	"log"
	"fmt"
)
func main() {
	e := Estudiante {
		Name: "Alejandro",
		Age: 30,
		Active: true
	}

	err := Crear(e)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create with exited")
}