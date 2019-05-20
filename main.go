package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	/*"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/go-chi/chi"
	*/
)
var note Note

// Punto de ejecución del ejecutable.
func main() {
	// Instancia de http.DefaultServeMux
	mux := http.NewServeMux()

	// flag para realizar la creación de las tablas en la base de datos.
	migrate := flag.Bool("migrate", false, "Crea las tablas en la base de datos")
	flag.Parse()

	if *migrate {
		if err := MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	// Rutas a manejar
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/notes", NotesHandler)

	// Log informativo
	log.Println("Corriendo en http://localhost:8081")

	// Servidor escuchando en el puerto 8080
	http.ListenAndServe(":8081", mux)
	
/*
	r := chi.NewRouter()
	r.Get("/public", func(w http.ResponseWriter, r *http.Request) {
		
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))

	http.ListenAndServe(":8081", r)
*/
}


// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
/*
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
*/

// IndexHandler nos permite manejar la petición a la ruta '/' y retornar "hola mundo"
// como respuesta al cliente.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hola mundo")
}


// GetNotesHandler nos permite manejar las peticiones a la ruta
// ‘/notes’ con el método GET.
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
    // Puntero a una estructura de tipo Note.
    n := new(Note)
    // Solicitando todas las notas en la base de datos.
    notes, err := n.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    // Convirtiendo el slice de notas a formato JSON,
    // retorna un []byte y un error.
    j, err := json.Marshal(notes)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Escribiendo el código de respuesta.
    w.WriteHeader(http.StatusOK)
    // Estableciendo el tipo de contenido del cuerpo de la
    // respuesta.
    w.Header().Set("Content-Type", "application/json")
    // Escribiendo la respuesta, es decir nuestro slice de notas
    // en formato JSON.
    w.Write(j)
}

// CreateNotesHandler nos permite manejar las peticiones a la ruta
// ‘/notes’ con el método POST.
func CreateNotesHandler(w http.ResponseWriter, r *http.Request) {
    var note Note
// Tomando el cuerpo de la petición, en formato JSON, y
    // decodificándola e la variable note que acabamos de
    // declarar.
    err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
// Creamos la nueva nota gracias al método Create.
    err = note.Create()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// UpdateNotesHandler nos permite manejar las peticiones a la ruta
// ‘/notes’ con el método UPDATE.
func UpdateNotesHandler(w http.ResponseWriter, r *http.Request) {
    var note Note
err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Actualizamos la nota correspondiente.
    err = note.Update()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// DeleteNotesHandler nos permite manejar las peticiones a la ruta
// ‘/notes’ con el método DELETE.
func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
    // obtenemos el valor pasado en la url como query
    // correspondiente a id, del tipo /notes?id=3.
    idStr := r.URL.Query().Get("id")
    // Verificamos que no esté vacío.
    if idStr == "" {
         http.Error(w, "Query id es requerido",
             http.StatusBadRequest)
         return
    }
    // Convertimos el valor obtenido del query a un int, de ser
    // posible.
    id, err := strconv.Atoi(idStr)
    if err != nil {
         http.Error(w, “Query id debe ser un número”,
             http.StatusBadRequest)
         return
    }
    var note Note
    // Borramos la nota con el id correspondiente.
    err = note.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// NotesHandler nos permite manejar la petición a la ruta ‘/notes’ // y pasa el control a la función correspondiente según el método
// de la petición.
func NotesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            GetNotesHandler(w, r)
        case http.MethodPost:
            CreateNotesHandler(w, r)
        case http.MethodPut:
            UpdateNotesHandler(w, r)
        case http.MethodDelete:
            DeleteNotesHandler(w, r)
        default:
            // Caso por defecto en caso de que se realice una
            // petición con un método diferente a los esperados.
            http.Error(w, "Metodo no permitido",
                http.StatusBadRequest)
            return
    }
}

