package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
	"errors"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
	/*"net/http" */
	"github.com/satori/go.uuid"
	 "os"
	"path/filepath"
	"strings"
	"github.com/go-chi/chi"
	// "bytes"
)
var Host string
var db *sql.DB
type UUID [16]byte
func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

		if _, err := db.Exec(
			`DROP TABLE IF EXISTS domain`); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(
			`CREATE TABLE IF NOT EXISTS domain (
				id INT PRIMARY KEY,
					title VARCHAR(64) NULL,
					description VARCHAR(200) NULL,
					Uuid VARCHAR(350),
					Host VARCHAR(120),
					Port INT,
					Protocol VARCHAR(120),
					IsPublic BOOL,
					Status   VARCHAR(80),
					StartTime       TIMESTAMP,
					TestTime        INT,
					EngineVersion   VARCHAR(120),
					CriteriaVersion VARCHAR(120),
					Endpoints       JSONB,
					Host__          JSONB
				)`); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(
			`DROP TABLE IF EXISTS accounts`); err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(
			`DROP TABLE IF EXISTS accountsafn`); err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(
			`DROP TABLE IF EXISTS estudiantes`); err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(
			`DROP TABLE IF EXISTS notes`); err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(
			`DROP TABLE IF EXISTS pruebaafn`); err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(
			`DROP TABLE IF EXISTS domaintest`); err != nil {
			log.Fatal(err)
		}
		
	
		if _, err := db.Exec(
			`CREATE TABLE IF NOT EXISTS domaintest (
					Uuid VARCHAR(350) NULL,
					Host VARCHAR(120) NULL,
					Port INT NULL,
					Protocol VARCHAR(120) NULL,
					IsPublic BOOL NULL,
					Status   VARCHAR(80) NULL,
					StartTime       DATE NULL,
					TestTime        INT NULL,
					EngineVersion   VARCHAR(120) NULL,
					CriteriaVersion VARCHAR(120) NULL,
					Endpoints       VARCHAR(8000) NULL,
					HostOld         VARCHAR(8000) NULL,
					HostNew         VARCHAR(8000) NULL
				)`); err != nil {
			log.Fatal(err)
		}
		// var t = time.Now().Unix()
		// fmt.Println(t)
		if _, err := db.Exec(
			`INSERT INTO domaintest (
					Uuid,
					Host,
					Port,
					Protocol, 
					IsPublic,
					Status,   
					StartTime,
					TestTime ,
					EngineVersion,   
					CriteriaVersion,
					endpoints,
					HostOld,
					HostNew
				) VALUES (
					'XXXX-YYYY-ZZZZ',
					'www.google.com',
					443,
					'http',
					false,
					'READY',
					'2019-03-26',
					1558624016,
					'1.34.2',
					'2009p',
					'{"endpoints": [
						{
						"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
						"serverName": "sfo03s08-in-x04.1e100.net",
						"statusMessage": "Ready",
						"grade": "A+",
						"gradeTrustIgnored": "A+",
						"hasWarnings": false,
						"isExceptional": true,
						"progress": 100,
						"duration": 85620,
						"delegation": 2
						},
						{
						"ipAddress": "172.217.6.36",
						"serverName": "sfo03s08-in-f4.1e100.net",
						"statusMessage": "Ready",
						"grade": "A+",
						"gradeTrustIgnored": "A+",
						"hasWarnings": false,
						"isExceptional": true,
						"progress": 100,
						"duration": 95185,
						"delegation": 2
						}
					  ]}',
					  '{"HostOld": [
						{
						"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
						"serverName": "sfo03s08-in-x04.1e100.net",
						"statusMessage": "Ready",
						"grade": "A+",
						"gradeTrustIgnored": "A+",
						"hasWarnings": false,
						"isExceptional": true,
						"progress": 100,
						"duration": 85620,
						"delegation": 2
						},
						{
						"ipAddress": "172.217.6.36",
						"serverName": "sfo03s08-in-f4.1e100.net",
						"statusMessage": "Ready",
						"grade": "A+",
						"gradeTrustIgnored": "A+",
						"hasWarnings": false,
						"isExceptional": true,
						"progress": 100,
						"duration": 95185,
						"delegation": 2
						}
					  ]}',
					  '{"HostNew": [
						{
						"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
						"serverName": "sfo03s08-in-x04.1e100.net",
						"statusMessage": "Ready",
						"grade": "A+",
						"gradeTrustIgnored": "A+",
						"hasWarnings": false,
						"isExceptional": true,
						"progress": 100,
						"duration": 85620,
						"delegation": 2
						},
						{
						"ipAddress": "172.217.6.36",
						"serverName": "sfo03s08-in-f4.1e100.net",
						"statusMessage": "Ready",
						"grade": "A+",
						"gradeTrustIgnored": "A+",
						"hasWarnings": false,
						"isExceptional": true,
						"progress": 100,
						"duration": 95185,
						"delegation": 2
						}
					  ]}'),
					  (
						'AAAA-BBBB-CCCC-DDDD',
						'www.googleafn.com',
						449,
						'http',
						false,
						'READY',
						'2019-03-26',
						1558624016,
						'1.34.2',
						'2009p',
						'{"endpoints": [
							{
							"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
							"serverName": "sfo03s08-in-x04.1e100.net",
							"statusMessage": "Ready",
							"grade": "A+",
							"gradeTrustIgnored": "A+",
							"hasWarnings": false,
							"isExceptional": true,
							"progress": 100,
							"duration": 85620,
							"delegation": 2
							},
							{
							"ipAddress": "172.217.6.36",
							"serverName": "sfo03s08-in-f4.1e100.net",
							"statusMessage": "Ready",
							"grade": "A+",
							"gradeTrustIgnored": "A+",
							"hasWarnings": false,
							"isExceptional": true,
							"progress": 100,
							"duration": 95185,
							"delegation": 2
							}
						  ]}',
						  '{"HostOld": [
							{
							"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
							"serverName": "sfo03s08-in-x04.1e100.net",
							"statusMessage": "Ready",
							"grade": "A+",
							"gradeTrustIgnored": "A+",
							"hasWarnings": false,
							"isExceptional": true,
							"progress": 100,
							"duration": 85620,
							"delegation": 2
							},
							{
							"ipAddress": "172.217.6.36",
							"serverName": "sfo03s08-in-f4.1e100.net",
							"statusMessage": "Ready",
							"grade": "A+",
							"gradeTrustIgnored": "A+",
							"hasWarnings": false,
							"isExceptional": true,
							"progress": 100,
							"duration": 95185,
							"delegation": 2
							}
						  ]}',
						  '{"HostNew": [
							{
							"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
							"serverName": "sfo03s08-in-x04.1e100.net",
							"statusMessage": "Ready",
							"grade": "A+",
							"gradeTrustIgnored": "A+",
							"hasWarnings": false,
							"isExceptional": true,
							"progress": 100,
							"duration": 85620,
							"delegation": 2
							},
							{
							"ipAddress": "172.217.6.36",
							"serverName": "sfo03s08-in-f4.1e100.net",
							"statusMessage": "Ready",
							"grade": "A+",
							"gradeTrustIgnored": "A+",
							"hasWarnings": false,
							"isExceptional": true,
							"progress": 100,
							"duration": 95185,
							"delegation": 2
							}
						  ]}')`); err != nil {
			log.Fatal(err)
		}
		 
 return db
}

func MakeMigrations() error {
	// GetConnection()

	db2 := GetConnection()
	defer db2.Close()

	q2 := `SELECT top 1 1 from domaintest`
	stmt2, err := db2.Prepare(q2)

	if err != nil {
		return err
	}
	defer stmt2.Close()

	r2, err := stmt2.Exec(q2)

	if err != nil {
		return err
	}

	i2, _ := r2.RowsAffected()
	fmt.Println("RowsAffected")
	fmt.Println(i2)

	// q := `CREATE TABLE IF NOT EXISTS notes (
	//         id INTEGER PRIMARY KEY AUTOINCREMENT,
    //    		title VARCHAR(64) NULL,
    //    		description VARCHAR(200) NULL
	//       );`

	// _, err := db.Exec(q)
	// if err != nil {
	// 	return err
	// }

	// q2 := `INSERT INTO 
	// 			notes(title, description)
	// 			VALUES ('aja','aja desc')`
			

	// db2 := GetConnection()
	// defer db2.Close()

	// stmt2, err := db2.Prepare(q2)

	// if err != nil {
	// 	return err
	// }
	// defer stmt2.Close()

	// r2, err := stmt2.Exec("ajaa", "descc")

	// if err != nil {
	// 	return err
	// }

	// i2, _ := r2.RowsAffected()

	// if i2 != 1 {
	// 	return errors.New("Should error rows i2")
	// }

	return nil
}


type Domaintest struct {
	Uuid            string      `json:"uuid"`
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	StartTime       int64       `json:"startTime"`
	TestTime        int64       `json:"testTime"`
	EngineVersion   string      `json:"engineVersion"`
	CriteriaVersion string      `json:"criteriaVersion"`
	Endpoints       []Endpoints `json:"endpoints"`
	HostOld         string `json:"HostOld"`
	HostNew         string `json:"HostNew"`
}
type DomaintestNuevos []Domaintest
type DomaintestOne struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	StartTime       int64       `json:"startTime"`
	TestTime        int64       `json:"testTime"`
	EngineVersion   string      `json:"engineVersion"`
	CriteriaVersion string      `json:"criteriaVersion"`
	Endpoints       []Endpoints `json:"endpoints"`
	HostOld         string `json:"HostOld"`
	HostNew         string `json:"HostNew"`
}
type DomaintestOnePruebas []DomaintestOne
type DomaintestOnes struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	StartTime       int64       `json:"startTime"`
	TestTime        int64       `json:"testTime"`
	EngineVersion   string      `json:"engineVersion"`
	CriteriaVersion string      `json:"criteriaVersion"`
	Endpoints       []Endpoints `json:"endpoints"`
	HostOld         string `json:"HostOld"`
	HostNew         string `json:"HostNew"`
}

type DomaintestsOnes []DomaintestOne

type Datas struct {
	nombre 		string `json:"nombre"`
}
type Note struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
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


type domain struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	StartTime       int64       `json:"startTime"`
	TestTime        int64       `json:"testTime"`
	EngineVersion   string      `json:"engineVersion"`
	CriteriaVersion string      `json:"criteriaVersion"`
	Endpoints       []Endpoints `json:"endpoints"`
	Host__            string      `json:"host__"`
	Port__            int         `json:"port__"`
	Protocol__        string      `json:"protocol__"`
	IsPublic__        bool        `json:"isPublic__"`
	Status__          string      `json:"status__"`
	StartTime__       int64       `json:"startTime__"`
	TestTime__        int64       `json:"testTime__"`
	EngineVersion__   string      `json:"engineVersion__"`
	CriteriaVersion__ string      `json:"criteriaVersion__"`
}
type Endpoints struct {
	IpAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int    `json:"duration"`
	Delegation        int    `json:"delegation"`
}

type Endpointss []Endpoints

func main() {

	GetConnection()
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	
	log.Println("Corriendo en http://localhost:8081")
	r := chi.NewRouter()

	r.Get("/public", func(w http.ResponseWriter, r *http.Request) {

		nombre := r.URL.Query().Get("nombre")
		
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		url := "https://api.ssllabs.com/api/v3/analyze?host="+nombre
		        
		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		j := "["+string(responseData)+"]"
		xp := []Domaintest{}
	
		errr := json.Unmarshal([]byte(j), &xp)
		if errr != nil {
			fmt.Println(errr)
		}
		data := Domaintest{}
		endpointsssf := Endpointss{}
		for i, v := range xp {
			 fmt.Println(i,v)
			 data.Host = v.Host
			 Host = v.Host
			 data.Port = v.Port
			 data.Protocol = v.Protocol
			 data.IsPublic = v.IsPublic
			 data.Status = v.Status
			 data.StartTime = v.StartTime
			 data.TestTime = v.TestTime
			 data.EngineVersion = v.EngineVersion
			 data.CriteriaVersion = v.CriteriaVersion
			
			for b, k := range v.Endpoints {
				endpointsss := Endpointss{
					Endpoints{
						Grade:k.Grade,
						IpAddress:k.IpAddress,
						ServerName : k.ServerName,
						StatusMessage : k.StatusMessage,
						GradeTrustIgnored : k.GradeTrustIgnored,
						HasWarnings : k.HasWarnings,
						IsExceptional : k.IsExceptional,
						Progress : k.Progress,
						Duration : k.Duration,
						Delegation : k.Delegation,
					},
				}
	
				fmt.Println(b)
				endpointsssf = endpointsss
			}
			data.Endpoints = endpointsssf 
			uuid, err := uuid.NewV4()
			fmt.Println(uuid)
			if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
			}
		}

		n := new(Domaintest)
		domain, err := n.GetAllDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		dataDomain, err := json.Marshal(domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		responsedata:= []Domaintest{}
	
		errrs := json.Unmarshal([]byte(dataDomain), &responsedata)
		if errrs != nil {
			fmt.Println(errrs)
		}

		fmt.Println(responsedata)
		if len(responsedata) > 0 {
			fmt.Println("debo de actualizar registros")
		} else {
			fmt.Println("debo de ingresar registro")
		}
		// fmt.Println(data)

		json.NewEncoder(w).Encode(data)
})
	// r.Get("/public", func (w http.ResponseWriter, r *http.Request){
	// 	nombre := r.URL.Query().Get("nombre")
	// 	fmt.Println(nombre)
	// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
	// 	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	url := "https://api.ssllabs.com/api/v3/analyze?host="+nombre
	// 	response, err := http.Get(url)

	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 		os.Exit(1)
	// 	}

	// 	responseData, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(responseData)
	// 	// j := "["+string(responseData)+"]"
			
	// 	// xp := []domain{}
	
	// 	// errr := json.Unmarshal([]byte(j), &xp)
		
	// 	// if errr != nil {
	// 	// 	fmt.Println(errr)
	// 	// }
	// 	//fmt.Println(len(responseData))
	// 	// fmt.Println("Inicio Entro por buscar metodo que se llama al dar clic en buscar:")
			
	// 	// for i, v := range xp {
	// 	// 	fmt.Printf("Something went wrong: %s", i)		
	// 	// 	uuid, err := uuid.NewV4()
	// 	// 	if err != nil {
	// 	// 	fmt.Printf("Something went wrong: %s", err)
	// 	// 	return
	// 	// 	}
			
	// 			//  newDomain[i].Host =  v.Host
	// 			// newDomain[i].Port = v.Port
	// 			// newDomain[i].Protocol = v.Protocol
	// 			// newDomain[i].IsPublic = v.IsPublic       
	// 			// newDomain[i].Status = v.Status      
	// 			// newDomain[i].StartTime = v.StartTime
	// 			// newDomain[i].TestTime = v.TestTime
	// 			// newDomain[i].EngineVersion = v.EngineVersion
	// 			// newDomain[i].CriteriaVersion = v.CriteriaVersion
			

	// 		// fmt.Println("entro en el metodo buscar debo guardar aca")
	// 		// fmt.Printf("Uuid")
	// 		// fmt.Printf("%s", uuid)
	// 		// fmt.Println("\t", v.Host)
	// 		// fmt.Println("\t", v.Port)
	// 		// fmt.Println("\t",v.Protocol)
	// 		// fmt.Println("\t",v.IsPublic)
	// 		// fmt.Println("\t",v.Status)
	// 		// fmt.Println("\t",v.StartTime)
	// 		// fmt.Println("\t",v.TestTime)
	// 		// fmt.Println("\t",v.EngineVersion)
	// 		// fmt.Println("\t",v.CriteriaVersion)
	// 		// fmt.Println("\t",v.Endpoints)
	// 		// for b, k := range v.Endpoints {
	// 		// 	fmt.Println("\t", uuid)
	// 		// 	fmt.Println("segundo recorrido")
	// 		// 	fmt.Println(b, k)
	// 		// 	fmt.Println("\t","IPAddress: " + string(k.IPAddress))
	// 		// 	fmt.Println("\t","ServerName: " + string(k.ServerName))
	// 		// 	fmt.Println("\t","StatusMessage: " + string(k.StatusMessage))
	// 		// 	fmt.Println("\t","Grade: " + string(k.Grade))
	// 		// 	fmt.Println("\t","GradeTrustIgnored: " + string(k.GradeTrustIgnored))
	// 		// 	fmt.Println("HasWarnings:\n",k.HasWarnings)
	// 		// 	fmt.Println("IsExceptional:\n",  k.IsExceptional)
	// 		// 	fmt.Println("Progress:\t", k.Progress)
	// 		// 	fmt.Println("Duration:\t",k.Duration)
	// 		// 	fmt.Println("Delegation:\t", + k.Delegation)

	// 			// newDomain.Endpoints = [
	// 			// 	"IPAddress" : k.IPAddress,
	// 			// 	"ServerName" : k.ServerName,
	// 			// 	"StatusMessage" : k.StatusMessage,
	// 			// 	"Grade" : k.Grade,
	// 			// 	"GradeTrustIgnored" : k.GradeTrustIgnored,
	// 			// 	"HasWarnings" : k.HasWarnings,
	// 			// 	"IsExceptional" : k.IsExceptional,
	// 			// 	"Progress" : k.Progress,
	// 			// 	"Duration" : k.Duration,
	// 			// 	"Delegation" : k.Delegation,
	// 			// ]					
	// 	// 	}
	// 	// }


	// 	n := new(Domaintest)
	// 	// Solicitando todas las notas en la base de datos.
	// 	domain, err := n.GetAllDomain()
	// 	fmt.Println("va a consultar bd")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 		fmt.Println("bummmm no trajo nada de bd")
	// 		return
	// 	}
	// 	// Convirtiendo el slice de notas a formato JSON,
	// 	// retorna un []byte y un error.
	// 	jj, err := json.Marshal(domain)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	// Escribiendo el código de respuesta.
	// 	w.WriteHeader(http.StatusOK)
	// 	// Estableciendo el tipo de contenido del cuerpo de la
	// 	// respuesta.
	// 	w.Header().Set("Content-Type", "application/json")
	// 	fmt.Println("devolver valor:")
	// 	fmt.Println(jj)
		
	// 	fmt.Println("Fin Entro por buscar metodo que se llama al dar clic en buscar:")
	// 	w.Write(responseData)
	// })
	// r.Get("/public", func(w http.ResponseWriter, r *http.Request) {
	// 		(w).Header().Set("Access-Control-Allow-Origin", "*")
	// 	    (w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 		nombre := r.URL.Query().Get("nombre")
			
	// 		if nombre == "" { 
	// 			nombre = "google.com"
	// 		  }
	// 		response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host="+nombre)

	// 		if err != nil {
	// 			fmt.Print(err.Error())
	// 			os.Exit(1)
	// 		}

	// 		responseData, err := ioutil.ReadAll(response.Body)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println("Inicio Entro por public metodo que se llama al inicio de el programa:")
	// 		fmt.Println("devolver data:")
	// 		fmt.Println(string(responseData))	
		
	// 		j := "["+string(responseData)+"]"
			
	// 		xp := []domain{}
		
	// 		errr := json.Unmarshal([]byte(j), &xp)
			
	// 		if errr != nil {
	// 			fmt.Println(errr)
	// 		}
	// 		fmt.Println(len(responseData))
	// 		for i, v := range xp {
	// 			fmt.Printf("Something went wrong: %s", i)		
	// 			uuid, err := uuid.NewV4()
	// 			if err != nil {
	// 			fmt.Printf("Something went wrong: %s", err)
	// 			return
	// 			}
				
	// 				//  newDomain[i].Host =  v.Host
	// 				// newDomain[i].Port = v.Port
	// 				// newDomain[i].Protocol = v.Protocol
	// 				// newDomain[i].IsPublic = v.IsPublic       
	// 				// newDomain[i].Status = v.Status      
	// 				// newDomain[i].StartTime = v.StartTime
	// 				// newDomain[i].TestTime = v.TestTime
	// 				// newDomain[i].EngineVersion = v.EngineVersion
	// 				// newDomain[i].CriteriaVersion = v.CriteriaVersion
				


	// 			//fmt.Println(i, v)
	// 			fmt.Printf("Uuid")
	// 			fmt.Printf("%s", uuid)
	// 			fmt.Println("\t", v.Host)
	// 			fmt.Println("\t", v.Port)
	// 			fmt.Println("\t",v.Protocol)
	// 			fmt.Println("\t",v.IsPublic)
	// 			fmt.Println("\t",v.Status)
	// 			fmt.Println("\t",v.StartTime)
	// 			fmt.Println("\t",v.TestTime)
	// 			fmt.Println("\t",v.EngineVersion)
	// 			fmt.Println("\t",v.CriteriaVersion)
	// 			fmt.Println("\t",v.Endpoints)
	// 			for b, k := range v.Endpoints {
	// 				fmt.Println("\t", uuid)
	// 				fmt.Println("segundo recorrido")
	// 				fmt.Println(b, k)
	// 				fmt.Println("\t","IPAddress: " + string(k.IPAddress))
	// 				fmt.Println("\t","ServerName: " + string(k.ServerName))
	// 				fmt.Println("\t","StatusMessage: " + string(k.StatusMessage))
	// 				fmt.Println("\t","Grade: " + string(k.Grade))
	// 				fmt.Println("\t","GradeTrustIgnored: " + string(k.GradeTrustIgnored))
	// 				fmt.Println("HasWarnings:\n",k.HasWarnings)
	// 				fmt.Println("IsExceptional:\n",  k.IsExceptional)
	// 				fmt.Println("Progress:\t", k.Progress)
	// 				fmt.Println("Duration:\t",k.Duration)
	// 				fmt.Println("Delegation:\t", + k.Delegation)

	// 				// newDomain.Endpoints = [
	// 				// 	"IPAddress" : k.IPAddress,
	// 				// 	"ServerName" : k.ServerName,
	// 				// 	"StatusMessage" : k.StatusMessage,
	// 				// 	"Grade" : k.Grade,
	// 				// 	"GradeTrustIgnored" : k.GradeTrustIgnored,
	// 				// 	"HasWarnings" : k.HasWarnings,
	// 				// 	"IsExceptional" : k.IsExceptional,
	// 				// 	"Progress" : k.Progress,
	// 				// 	"Duration" : k.Duration,
	// 				// 	"Delegation" : k.Delegation,
	// 				// ]					
	// 			}
	// 		}
	// 		// for key, result := range results {

	// 		// 	fmt.Println("Reading Value for Key :", key)
				
	// 		// 	fmt.Println("Id :", result["port"])
	// 		// }
	// 		fmt.Println("Fin Entro por public metodo que se llama al inicio de el programa:")
			
	// 		w.Write(responseData)
	// })


	
	r.Get("/buscartest", func(w http.ResponseWriter, r *http.Request) {

		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		
		n := new(Domaintest)
		// Solicitando todas las notas en la base de datos.
		domain, err := n.GetAllDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Convirtiendo el slice de notas a formato JSON,
		// retorna un []byte y un error.
		j, err := json.Marshal(domain)
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
	
})


	


		r.Get("/buscar", func(w http.ResponseWriter, r *http.Request) {

			nombre := r.URL.Query().Get("nombre")
			fmt.Println(nombre)
			(w).Header().Set("Access-Control-Allow-Origin", "*")
			(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			url := "https://api.ssllabs.com/api/v3/analyze?host="+nombre
			response, err := http.Get(url)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(responseData))	
			fmt.Println("paso por aca y funciona buscando debe recorrer para poder editar")
		
		
			fmt.Println(len(responseData))
		
			j := "["+string(responseData)+"]"
			//fmt.Println("json:",j)
			
			xp := []domain{}
			
			errr := json.Unmarshal([]byte(j), &xp)
			
			if errr != nil {
				fmt.Println(errr)
			}
			// fmt.Printf("go data: %+v\n", xp)
			fmt.Println(len(responseData))
			// var newDomain[len(responseData)] domain
			for i, v := range xp {
				fmt.Printf("afnSomething went wrong: %s", i)		
				// or error handling
				uuid, err := uuid.NewV4()
				if err != nil {
				fmt.Printf("Something went wrong: %s", err)
				return
				}
				
					//  newDomain[i].Host =  v.Host
					// newDomain[i].Port = v.Port
					// newDomain[i].Protocol = v.Protocol
					// newDomain[i].IsPublic = v.IsPublic       
					// newDomain[i].Status = v.Status      
					// newDomain[i].StartTime = v.StartTime
					// newDomain[i].TestTime = v.TestTime
					// newDomain[i].EngineVersion = v.EngineVersion
					// newDomain[i].CriteriaVersion = v.CriteriaVersion
				


				//fmt.Println(i, v)
				fmt.Printf("Uuid")
				fmt.Printf("%s", uuid)
				fmt.Println("\t", v.Host)
				fmt.Println("\t", v.Port)
				fmt.Println("\t",v.Protocol)
				fmt.Println("\t",v.IsPublic)
				fmt.Println("\t",v.Status)
				fmt.Println("\t",v.StartTime)
				fmt.Println("\t",v.TestTime)
				fmt.Println("\t",v.EngineVersion)
				fmt.Println("\t",v.CriteriaVersion)
				fmt.Println("\t",v.Endpoints)
				for b, k := range v.Endpoints {
					fmt.Println("\t", uuid)
					fmt.Println("segundo recorrido")
					fmt.Println(b, k)
					fmt.Println("\t","IPAddress: " + string(k.IpAddress))
					fmt.Println("\t","ServerName: " + string(k.ServerName))
					fmt.Println("\t","StatusMessage: " + string(k.StatusMessage))
					fmt.Println("\t","Grade: " + string(k.Grade))
					fmt.Println("\t","GradeTrustIgnored: " + string(k.GradeTrustIgnored))
					fmt.Println("HasWarnings:\n",k.HasWarnings)
					fmt.Println("IsExceptional:\n",  k.IsExceptional)
					fmt.Println("Progress:\t", k.Progress)
					fmt.Println("Duration:\t",k.Duration)
					fmt.Println("Delegation:\t", + k.Delegation)

					// newDomain.Endpoints = [
					// 	"IPAddress" : k.IPAddress,
					// 	"ServerName" : k.ServerName,
					// 	"StatusMessage" : k.StatusMessage,
					// 	"Grade" : k.Grade,
					// 	"GradeTrustIgnored" : k.GradeTrustIgnored,
					// 	"HasWarnings" : k.HasWarnings,
					// 	"IsExceptional" : k.IsExceptional,
					// 	"Progress" : k.Progress,
					// 	"Duration" : k.Duration,
					// 	"Delegation" : k.Delegation,
					// ]					
				}
			}
			// for key, result := range results {

			// 	fmt.Println("Reading Value for Key :", key)
				
			// 	fmt.Println("Id :", result["port"])
			// }
			//  w.Write(j)
		
			 w.Write(responseData) 
	})
	
// 	r.Get("/buscar", func(w http.ResponseWriter, r *http.Request) {

// 		nombre := r.URL.Query().Get("nombre")
// 		fmt.Println(nombre)
// 		(w).Header().Set("Access-Control-Allow-Origin", "*")
// 		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		url := "https://api.ssllabs.com/api/v3/analyze?host="+nombre
// 		response, err := http.Get(url)

// 		if err != nil {
// 			fmt.Print(err.Error())
// 			os.Exit(1)
// 		}

// 		responseData, err := ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
		
// 		// j := "["+string(responseData)+"]"
// 		// xp := []domain{}
	
// 		// errr := json.Unmarshal([]byte(j), &xp)
		
// 		// if errr != nil {
// 		// 	fmt.Println(errr)
// 		// }
// 		// fmt.Println(string(responseData))
// 		// for i, v := range xp {
// 		// 	fmt.Println("le id:",i)
// 		// 	var datoAfn = v.Host
			
// 		// 	fmt.Println("este valor lo devolvió la data", datoAfn)
// 		// 	fmt.Println("este : es: es:", datoAfn)
// 		// }

// 		j := "["+string(responseData)+"]"
	
// 		xp := []domain{}
	
// 		errr := json.Unmarshal([]byte(j), &xp)
		
// 		if errr != nil {
// 			fmt.Println(errr)
// 		}
// 		//fmt.Println(len(responseData))
// 		fmt.Println("Inicio Entro por buscar metodo que se llama al dar clic en buscar:")
			
// 		for i, v := range xp {
// 			fmt.Printf("Something went wrong: %s", i)		
// 			uuid, err := uuid.NewV4()
// 			if err != nil {
// 			fmt.Printf("Something went wrong: %s", err)
// 			return
// 			}
			
// 				//  newDomain[i].Host =  v.Host
// 				// newDomain[i].Port = v.Port
// 				// newDomain[i].Protocol = v.Protocol
// 				// newDomain[i].IsPublic = v.IsPublic       
// 				// newDomain[i].Status = v.Status      
// 				// newDomain[i].StartTime = v.StartTime
// 				// newDomain[i].TestTime = v.TestTime
// 				// newDomain[i].EngineVersion = v.EngineVersion
// 				// newDomain[i].CriteriaVersion = v.CriteriaVersion
			

// 			fmt.Println("entro en el metodo buscar debo guardar aca")
// 			//fmt.Println(i, v)
// 			fmt.Printf("Uuid")
// 			fmt.Printf("%s", uuid)
// 			fmt.Println("\t", v.Host)
// 			fmt.Println("\t", v.Port)
// 			fmt.Println("\t",v.Protocol)
// 			fmt.Println("\t",v.IsPublic)
// 			fmt.Println("\t",v.Status)
// 			fmt.Println("\t",v.StartTime)
// 			fmt.Println("\t",v.TestTime)
// 			fmt.Println("\t",v.EngineVersion)
// 			fmt.Println("\t",v.CriteriaVersion)
// 			fmt.Println("\t",v.Endpoints)
// 			for b, k := range v.Endpoints {
// 				fmt.Println("\t", uuid)
// 				fmt.Println("segundo recorrido")
// 				fmt.Println(b, k)
// 				fmt.Println("\t","IPAddress: " + string(k.IPAddress))
// 				fmt.Println("\t","ServerName: " + string(k.ServerName))
// 				fmt.Println("\t","StatusMessage: " + string(k.StatusMessage))
// 				fmt.Println("\t","Grade: " + string(k.Grade))
// 				fmt.Println("\t","GradeTrustIgnored: " + string(k.GradeTrustIgnored))
// 				fmt.Println("HasWarnings:\n",k.HasWarnings)
// 				fmt.Println("IsExceptional:\n",  k.IsExceptional)
// 				fmt.Println("Progress:\t", k.Progress)
// 				fmt.Println("Duration:\t",k.Duration)
// 				fmt.Println("Delegation:\t", + k.Delegation)

// 				// newDomain.Endpoints = [
// 				// 	"IPAddress" : k.IPAddress,
// 				// 	"ServerName" : k.ServerName,
// 				// 	"StatusMessage" : k.StatusMessage,
// 				// 	"Grade" : k.Grade,
// 				// 	"GradeTrustIgnored" : k.GradeTrustIgnored,
// 				// 	"HasWarnings" : k.HasWarnings,
// 				// 	"IsExceptional" : k.IsExceptional,
// 				// 	"Progress" : k.Progress,
// 				// 	"Duration" : k.Duration,
// 				// 	"Delegation" : k.Delegation,
// 				// ]					
// 			}
// 		}


// 		n := new(Domaintest)
// 		// Solicitando todas las notas en la base de datos.
// 		domain, err := n.GetAllDomain()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusNotFound)
// 			return
// 		}
// 		// Convirtiendo el slice de notas a formato JSON,
// 		// retorna un []byte y un error.
// 		// j := "["+string(domain)+"]"
// 		jj, err := json.Marshal(domain)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		// Escribiendo el código de respuesta.
// 		w.WriteHeader(http.StatusOK)
// 		// Estableciendo el tipo de contenido del cuerpo de la
// 		// respuesta.
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Println("devolver valor:")
// 		fmt.Println(jj)
		
// 		fmt.Println("Fin Entro por buscar metodo que se llama al dar clic en buscar:")
// 		w.Write(jj)
// })

	r.Get("/notes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
	         GetNotesHandler(w, r)
	      default:
            http.Error(w, "Metodo no permitido",
                http.StatusBadRequest)
            return
    }
	})
	
	r.Post("/notes", func(w http.ResponseWriter, r *http.Request) {
          CreateNotesHandler(w, r)
	})

	r.Put("/notes", func(w http.ResponseWriter, r *http.Request) {
		UpdateNotesHandler(w, r)
	})

	r.Delete("/notes", func(w http.ResponseWriter, r *http.Request) {
		DeleteNotesHandler(w, r)
	})


	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))

	http.ListenAndServe(":8081", r)

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
	router.Post("/notes", CreateNotesHandler)
	router.Put("/notes/{id}",UpdateNotesHandler)
	router.Delete("/notes/{id}",DeleteNotesHandler)

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

func (n *Domaintest) GetAllDomain() ([]Domaintest, error) {
	db := GetConnection()
	Host = "'"+Host+"'"

	q := "SELECT distinct uuid,host,port FROM domaintest where host="+string(Host)
	// Ejecutamos la query
	rows, err := db.Query(q)
	if err != nil {
		return []Domaintest{}, err
	}
	defer rows.Close()
	bks := make([]Domaintest, 0)
	for rows.Next() {
		bk := Domaintest{}
		err := rows.Scan(&bk.Uuid,&bk.Host, &bk.Port) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	return bks, nil 
	
    // fmt.Println("dentro de la consulta de bd rows ")
	// fmt.Println(rows)
	// for rows.Next() {
	// 	rows.Scan(&n.Host,&n.Uuid)
	// 	domain = append(domain, *n)
	// }
	// fmt.Println("dentro de la consulta de bd ")
	// fmt.Println(domain)
	// fmt.Println("termino la consulta de bd dentro")
	// return domain, nil
}