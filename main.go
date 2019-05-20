package main

import (
	"encoding/json"
	/*"flag"*/
	"fmt"
	"log"
	"net/http"
	"strconv"

	"errors"
	"time"
	/*_ "github.com/mattn/go-sqlite3"*/

	"database/sql"
	_ "github.com/lib/pq"
	/*"net/http"*/
	 "os"
	"path/filepath"
	"strings"
	
	
	"github.com/go-chi/chi"
	
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	/*
	db, err = sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}
	return db*/

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	//db, err := sql.Open("postgres", dsn)

	    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    // Create the "accounts" table.
    if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS accountsafn (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
		}
	
		if _, err := db.Exec(
			`CREATE TABLE IF NOT EXISTS notes (
				id INT PRIMARY KEY,
				   title VARCHAR(64) NULL,
				   description VARCHAR(200) NULL
			  )`); err != nil {
			log.Fatal(err)
		}
 
 return db
}

func MakeMigrations() error {
	db := GetConnection()
	q := `CREATE TABLE IF NOT EXISTS notes (
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
       		title VARCHAR(64) NULL,
       		description VARCHAR(200) NULL
	      );`

	_, err := db.Exec(q)
	if err != nil {
		return err
	}

	q2 := `INSERT INTO 
				notes(title, description)
				VALUES ('aja','aja desc')`
			

	db2 := GetConnection()
	defer db2.Close()

	stmt2, err := db2.Prepare(q2)

	if err != nil {
		return err
	}
	defer stmt2.Close()

	r2, err := stmt2.Exec("ajaa", "descc")

	if err != nil {
		return err
	}

	i2, _ := r2.RowsAffected()

	if i2 != 1 {
		return errors.New("Should error rows i2")
	}

	return nil
}


type Note struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func (n Note) Create() error {
	// Realizamos la conexión a la base de datos.
	db := GetConnection()

	// Query para insertar los datos en la tabla notes
	q := `INSERT INTO notes (id,title, description)
			VALUES($1, $2, $3)`

	// Preparamos la petición para insertar los datos de manera segura
	// y evitar código malicioso.
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Ejecutamos la petición pasando los datos correspondientes. El orden
	// es importante, corresponde con los "?" delstring q.
	r, err := stmt.Exec(n.ID,n.Title, n.Description)
	if err != nil {
		return err
	}
	// Confirmamos que una fila fuera afectada, debido a que insertamos un
	// registro en la tabla. En caso contrario devolvemos un error.
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	// Si llegamos a este punto consideramos que todo el proceso fue exitoso
	// y retornamos un nil para confirmar que no existe un error.
	return nil
}

func (n *Note) GetAll() ([]Note, error) {
	db := GetConnection()
	q := `SELECT
			id, title, description
			FROM notes`
	// Ejecutamos la query
	rows, err := db.Query(q)
	if err != nil {
		return []Note{}, err
	}
	// Cerramos el recurso
	defer rows.Close()

	// Declaramos un slice de notas para que almacene las notas que retorne
	// la petición.
	notes := []Note{}
	// El método Next retorna un bool, mientras sea true indicará que existe
	// un valor siguiente para leer.
	for rows.Next() {
		// Escaneamos el valor actual de la fila e insertamos el retorno
		// en los correspondientes campos de la nota.
		rows.Scan(&n.ID, &n.Title, &n.Description)
		// Añadimos cada nueva nota al slice de notas que declaramos antes.
		notes = append(notes, *n)
	}
	return notes, nil
}

func (n *Note) GetByID(id int) (Note, error) {
	db := GetConnection()
	q := `SELECT
		id, title, description
		FROM notes WHERE id=$1`

	err := db.QueryRow(q, id).Scan(
		&n.ID, &n.Title, &n.Description,
	)
	if err != nil {
		return Note{}, err
	}

	return *n, nil
}

func (n Note) Update() error {
	db := GetConnection()
	q := `UPDATE notes set title=$1, description=$2
		WHERE id=$3`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(n.Title, n.Description, n.ID)
	if err != nil {
		return err
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	return nil
}

func (n Note) Delete(id int) error {
	db := GetConnection()

	q := `DELETE FROM notes
		WHERE id=$1`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	return nil
}



/*var note Note*/

// Punto de ejecución del ejecutable.
func main() {
	
	// Instancia de http.DefaultServeMux
	mux := http.NewServeMux()

	// flag para realizar la creación de las tablas en la base de datos.
	

	mux.HandleFunc("/", IndexHandler)
	//mux.HandleFunc("/notes", NotesHandler)

	log.Println("Corriendo en http://localhost:8081")

	//http.ListenAndServe(":8081", mux)
	r := chi.NewRouter()
	r.Get("/public", func(w http.ResponseWriter, r *http.Request) {
		
	})

	r.Get("/notes", func(w http.ResponseWriter, r *http.Request) {
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
	})	


	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))

	http.ListenAndServe(":8081", r)

	/*
	r := chi.NewRouter()
	r.Get("/public", func(w http.ResponseWriter, r *http.Request) {
		
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))

	http.ListenAndServe(":8081", r)
*/
/*
	router = chi.NewRouter() 
	http.ListenAndServe(":8081", Logger())*/
}

func Logger() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(time.Now(), r.Method, r.URL)
        router.ServeHTTP(w, r) // dispatch the request
    })
}

var router *chi.Mux
func routers() *chi.Mux {
    router.Get("/notes", GetNotesHandler)
 /*   router.Get("/posts/{id}", DetailPost)
    router.Post("/posts", CreatePost)
    router.Put("/posts/{id}", UpdatePost)
    router.Delete("/posts/{id}", DeletePost)
   */ 
    return router
}


// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.

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


// IndexHandler nos permite manejar la petición a la ruta '/' y retornar "hola mundo"
// como respuesta al cliente.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprint(w, "hola mundo")
	direccion := ":8081" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion+"/public/index.html", nil))
	
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
         http.Error(w, "Query id debe ser un número",
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

