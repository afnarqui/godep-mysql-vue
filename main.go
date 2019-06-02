package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	 "os"
	"path/filepath"
	"strings"
	"github.com/go-chi/chi"
	"errors"
)
var Host string
var db *sql.DB
type UUID [16]byte
var domainnew = Domain{}

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

		// if _, err := db.Exec(
		// 	`DROP TABLE IF EXISTS domain`); err != nil {
		// 	log.Fatal(err)
		// }

		// if _, err := db.Exec(
		// 	`CREATE TABLE IF NOT EXISTS domain (
		// 		id INT PRIMARY KEY,
		// 			title VARCHAR(64) NULL,
		// 			description VARCHAR(200) NULL,
		// 			Uuid VARCHAR(350),
		// 			Host VARCHAR(120),
		// 			Port INT,
		// 			Protocol VARCHAR(120),
		// 			IsPublic BOOL,
		// 			Status   VARCHAR(80),
		// 			StartTime       TIMESTAMP,
		// 			TestTime        INT,
		// 			EngineVersion   VARCHAR(120),
		// 			CriteriaVersion VARCHAR(120),
		// 			Endpoints       JSONB,
		// 			Host__          JSONB
		// 		)`); err != nil {
		// 	log.Fatal(err)
		// }


		// if _, err := db.Exec(
		// 	`DROP TABLE IF EXISTS domainold`); err != nil {
		// 	log.Fatal(err)
		// }
		
		// if _, err := db.Exec(
		// 	`CREATE TABLE IF NOT EXISTS domain (
		// 			Uuid uuid NULL,
		// 			Host VARCHAR(120) NULL,
		// 			Port INT NULL,
		// 			Protocol VARCHAR(120) NULL,
		// 			IsPublic BOOL NULL,
		// 			Status   VARCHAR(80) NULL,
		// 			StartTime       DATE NULL,
		// 			TestTime        INT NULL,
		// 			EngineVersion   VARCHAR(120) NULL,
		// 			CriteriaVersion VARCHAR(120) NULL,
		// 			Endpoints       VARCHAR(8000) NULL,
		// 			HostOld         VARCHAR(8000) NULL,
		// 			HostNew         VARCHAR(8000) NULL
		// 		)`); err != nil {
		// 	log.Fatal(err)
		// }

		// if _, err := db.Exec(
		// 	`CREATE TABLE IF NOT EXISTS domainhistory (
		// 			Uuid uuid NULL,
		// 			Host VARCHAR(120) NULL,
		// 			Port INT NULL,
		// 			Protocol VARCHAR(120) NULL,
		// 			IsPublic BOOL NULL,
		// 			Status   VARCHAR(80) NULL,
		// 			StartTime       DATE NULL,
		// 			TestTime        INT NULL,
		// 			EngineVersion   VARCHAR(120) NULL,
		// 			CriteriaVersion VARCHAR(120) NULL,
		// 			Endpoints       VARCHAR(8000) NULL,
		// 			HostOld         VARCHAR(8000) NULL,
		// 			HostNew         VARCHAR(8000) NULL
		// 		)`); err != nil {
		// 	log.Fatal(err)
		// }

		// if _, err := db.Exec(
		// 	`CREATE TABLE IF NOT EXISTS domain (
		// 			Uuid uuid NULL,
		// 			Host VARCHAR(120) NULL,
		// 			Port INT NULL,
		// 			Protocol VARCHAR(120) NULL,
		// 			IsPublic BOOL NULL,
		// 			Status   VARCHAR(80) NULL,
		// 			StartTime       DATE NULL,
		// 			TestTime        INT NULL,
		// 			EngineVersion   VARCHAR(120) NULL,
		// 			CriteriaVersion VARCHAR(120) NULL,
		// 			Endpoints       VARCHAR(8000) NULL,
		// 			HostOld         VARCHAR(8000) NULL,
		// 			HostNew         VARCHAR(8000) NULL
		// 		)`); err != nil {
		// 	log.Fatal(err)
		// }

		// if _, err := db.Exec(
		// 	`INSERT INTO domain (
		// 			Host,
		// 			Port,
		// 			Protocol, 
		// 			IsPublic,
		// 			Status,   
		// 			StartTime,
		// 			TestTime ,
		// 			EngineVersion,   
		// 			CriteriaVersion,
		// 			endpoints,
		// 			HostOld,
		// 			HostNew
		// 		) VALUES (
		// 			'www.google.com',
		// 			443,
		// 			'http',
		// 			false,
		// 			'READY',
		// 			'2019-03-26',
		// 			1558624016,
		// 			'1.34.2',
		// 			'2009p',
		// 			'{"endpoints": [
		// 				{
		// 				"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
		// 				"serverName": "sfo03s08-in-x04.1e100.net",
		// 				"statusMessage": "Ready",
		// 				"grade": "A+",
		// 				"gradeTrustIgnored": "A+",
		// 				"hasWarnings": false,
		// 				"isExceptional": true,
		// 				"progress": 100,
		// 				"duration": 85620,
		// 				"delegation": 2
		// 				},
		// 				{
		// 				"ipAddress": "172.217.6.36",
		// 				"serverName": "sfo03s08-in-f4.1e100.net",
		// 				"statusMessage": "Ready",
		// 				"grade": "A+",
		// 				"gradeTrustIgnored": "A+",
		// 				"hasWarnings": false,
		// 				"isExceptional": true,
		// 				"progress": 100,
		// 				"duration": 95185,
		// 				"delegation": 2
		// 				}
		// 			  ]}',
		// 			  '{"HostOld": [
		// 				{
		// 				"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
		// 				"serverName": "sfo03s08-in-x04.1e100.net",
		// 				"statusMessage": "Ready",
		// 				"grade": "A+",
		// 				"gradeTrustIgnored": "A+",
		// 				"hasWarnings": false,
		// 				"isExceptional": true,
		// 				"progress": 100,
		// 				"duration": 85620,
		// 				"delegation": 2
		// 				},
		// 				{
		// 				"ipAddress": "172.217.6.36",
		// 				"serverName": "sfo03s08-in-f4.1e100.net",
		// 				"statusMessage": "Ready",
		// 				"grade": "A+",
		// 				"gradeTrustIgnored": "A+",
		// 				"hasWarnings": false,
		// 				"isExceptional": true,
		// 				"progress": 100,
		// 				"duration": 95185,
		// 				"delegation": 2
		// 				}
		// 			  ]}',
		// 			  '{"HostNew": [
		// 				{
		// 				"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
		// 				"serverName": "sfo03s08-in-x04.1e100.net",
		// 				"statusMessage": "Ready",
		// 				"grade": "A+",
		// 				"gradeTrustIgnored": "A+",
		// 				"hasWarnings": false,
		// 				"isExceptional": true,
		// 				"progress": 100,
		// 				"duration": 85620,
		// 				"delegation": 2
		// 				},
		// 				{
		// 				"ipAddress": "172.217.6.36",
		// 				"serverName": "sfo03s08-in-f4.1e100.net",
		// 				"statusMessage": "Ready",
		// 				"grade": "A+",
		// 				"gradeTrustIgnored": "A+",
		// 				"hasWarnings": false,
		// 				"isExceptional": true,
		// 				"progress": 100,
		// 				"duration": 95185,
		// 				"delegation": 2
		// 				}
		// 			  ]}'),
		// 			  (
		// 				'www.googleafn.com',
		// 				449,
		// 				'http',
		// 				false,
		// 				'READY',
		// 				'2019-03-26',
		// 				1558624016,
		// 				'1.34.2',
		// 				'2009p',
		// 				'{"endpoints": [
		// 					{
		// 					"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
		// 					"serverName": "sfo03s08-in-x04.1e100.net",
		// 					"statusMessage": "Ready",
		// 					"grade": "A+",
		// 					"gradeTrustIgnored": "A+",
		// 					"hasWarnings": false,
		// 					"isExceptional": true,
		// 					"progress": 100,
		// 					"duration": 85620,
		// 					"delegation": 2
		// 					},
		// 					{
		// 					"ipAddress": "172.217.6.36",
		// 					"serverName": "sfo03s08-in-f4.1e100.net",
		// 					"statusMessage": "Ready",
		// 					"grade": "A+",
		// 					"gradeTrustIgnored": "A+",
		// 					"hasWarnings": false,
		// 					"isExceptional": true,
		// 					"progress": 100,
		// 					"duration": 95185,
		// 					"delegation": 2
		// 					}
		// 				  ]}',
		// 				  '{"HostOld": [
		// 					{
		// 					"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
		// 					"serverName": "sfo03s08-in-x04.1e100.net",
		// 					"statusMessage": "Ready",
		// 					"grade": "A+",
		// 					"gradeTrustIgnored": "A+",
		// 					"hasWarnings": false,
		// 					"isExceptional": true,
		// 					"progress": 100,
		// 					"duration": 85620,
		// 					"delegation": 2
		// 					},
		// 					{
		// 					"ipAddress": "172.217.6.36",
		// 					"serverName": "sfo03s08-in-f4.1e100.net",
		// 					"statusMessage": "Ready",
		// 					"grade": "A+",
		// 					"gradeTrustIgnored": "A+",
		// 					"hasWarnings": false,
		// 					"isExceptional": true,
		// 					"progress": 100,
		// 					"duration": 95185,
		// 					"delegation": 2
		// 					}
		// 				  ]}',
		// 				  '{"HostNew": [
		// 					{
		// 					"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
		// 					"serverName": "sfo03s08-in-x04.1e100.net",
		// 					"statusMessage": "Ready",
		// 					"grade": "A+",
		// 					"gradeTrustIgnored": "A+",
		// 					"hasWarnings": false,
		// 					"isExceptional": true,
		// 					"progress": 100,
		// 					"duration": 85620,
		// 					"delegation": 2
		// 					},
		// 					{
		// 					"ipAddress": "172.217.6.36",
		// 					"serverName": "sfo03s08-in-f4.1e100.net",
		// 					"statusMessage": "Ready",
		// 					"grade": "A+",
		// 					"gradeTrustIgnored": "A+",
		// 					"hasWarnings": false,
		// 					"isExceptional": true,
		// 					"progress": 100,
		// 					"duration": 95185,
		// 					"delegation": 2
		// 					}
		// 				  ]}')`); err != nil {
		// 	log.Fatal(err)
		// }
		 
 return db
}

func (n *Domain) GetAllDomain() ([]Domain, error) {
	db := GetConnection()
	Host = "'"+Host+"'"

	q := "SELECT distinct uuid,host,port FROM domain where host="+string(Host)
	rows, err := db.Query(q)
	if err != nil {
		return []Domain{}, err
	}
	defer rows.Close()
	bks := make([]Domain, 0)
	for rows.Next() {
		bk := Domain{}
		err := rows.Scan(&bk.Uuid,&bk.Host, &bk.Port) 
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	return bks, nil 
}

type Domain struct {
	Uuid            string 		`json:"uuid"`
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
		xp := []Domain{}
	
		errr := json.Unmarshal([]byte(j), &xp)
		if errr != nil {
			fmt.Println(errr)
		}
		data := Domain{}
		endpointsssf := Endpointss{}

		for i, v := range xp {
			uuid, err := uuid.NewV4()
			fmt.Println(uuid)
			if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
			}
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
		}

		n := new(Domain)
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

		responsedata:= []Domain{}
	
		errrs := json.Unmarshal([]byte(dataDomain), &responsedata)
		if errrs != nil {
			fmt.Println(errrs)
		}

		fmt.Println(responsedata)
		if len(responsedata) > 0 {
			fmt.Println("debo de actualizar registros")
		} else {
			var datanew Domain
			fmt.Println("debo de ingresar registro")
			domainnew = data
			fmt.Println(domainnew)
			err = datanew.CreateDomain()
		}

		json.NewEncoder(w).Encode(data)
})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))

	http.ListenAndServe(":8081", r)

}

func Logger() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(time.Now(), r.Method, r.URL)
        router.ServeHTTP(w, r) 
    })
}
var router *chi.Mux

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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprint(w, "hola mundo")
	direccion := ":8081" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)

	log.Fatal(http.ListenAndServe(direccion+"/public/index.html", nil))
	
}

func (n Domain) CreateDomain() error {
	// db := GetConnection()

	// q := `INSERT INTO domain (
	// 					Host,
	// 					Port,
	// 					Protocol, 
	// 					IsPublic,
	// 					Status,   
	// 					StartTime,
	// 					TestTime ,
	// 					EngineVersion,   
	// 					CriteriaVersion,
	// 					endpoints,
	// 					HostOld,
	// 					HostNew
	// 				) VALUES (
	// 					'www.google.com',
	// 					443,
	// 					'http',
	// 					false,
	// 					'READY',
	// 					'2019-03-26',
	// 					1558624016,
	// 					'1.34.2',
	// 					'2009p',
	// 					'{"endpoints": [
	// 						{
	// 						"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
	// 						"serverName": "sfo03s08-in-x04.1e100.net",
	// 						"statusMessage": "Ready",
	// 						"grade": "A+",
	// 						"gradeTrustIgnored": "A+",
	// 						"hasWarnings": false,
	// 						"isExceptional": true,
	// 						"progress": 100,
	// 						"duration": 85620,
	// 						"delegation": 2
	// 						},
	// 						{
	// 						"ipAddress": "172.217.6.36",
	// 						"serverName": "sfo03s08-in-f4.1e100.net",
	// 						"statusMessage": "Ready",
	// 						"grade": "A+",
	// 						"gradeTrustIgnored": "A+",
	// 						"hasWarnings": false,
	// 						"isExceptional": true,
	// 						"progress": 100,
	// 						"duration": 95185,
	// 						"delegation": 2
	// 						}
	// 					  ]}',
	// 					  '{"HostOld": [
	// 						{
	// 						"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
	// 						"serverName": "sfo03s08-in-x04.1e100.net",
	// 						"statusMessage": "Ready",
	// 						"grade": "A+",
	// 						"gradeTrustIgnored": "A+",
	// 						"hasWarnings": false,
	// 						"isExceptional": true,
	// 						"progress": 100,
	// 						"duration": 85620,
	// 						"delegation": 2
	// 						},
	// 						{
	// 						"ipAddress": "172.217.6.36",
	// 						"serverName": "sfo03s08-in-f4.1e100.net",
	// 						"statusMessage": "Ready",
	// 						"grade": "A+",
	// 						"gradeTrustIgnored": "A+",
	// 						"hasWarnings": false,
	// 						"isExceptional": true,
	// 						"progress": 100,
	// 						"duration": 95185,
	// 						"delegation": 2
	// 						}
	// 					  ]}',
	// 					  '{"HostNew": [
	// 						{
	// 						"ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
	// 						"serverName": "sfo03s08-in-x04.1e100.net",
	// 						"statusMessage": "Ready",
	// 						"grade": "A+",
	// 						"gradeTrustIgnored": "A+",
	// 						"hasWarnings": false,
	// 						"isExceptional": true,
	// 						"progress": 100,
	// 						"duration": 85620,
	// 						"delegation": 2
	// 						},
	// 						{
	// 						"ipAddress": "172.217.6.36",
	// 						"serverName": "sfo03s08-in-f4.1e100.net",
	// 						"statusMessage": "Ready",
	// 						"grade": "A+",
	// 						"gradeTrustIgnored": "A+",
	// 						"hasWarnings": false,
	// 						"isExceptional": true,
	// 						"progress": 100,
	// 						"duration": 95185,
	// 						"delegation": 2
	// 						}
	// 					  ]}')`

	var host = domainnew.Host
	var port = domainnew.Port
	fmt.Println(host)
	
	fmt.Println(port)
	// if _, err := db.Exec("INSERT INTO domain (Host,Port) VALUES ($host,$port)"); err != nil {
	// 	log.Fatal(err)
	// }

	q := `INSERT INTO 
	domain(host,port)
	VALUES ($1,$2)`


		db := GetConnection()
		defer db.Close()

		stmt, err := db.Prepare(q)

		if err != nil {
		return err
		}
		defer stmt.Close()

		r, err := stmt.Exec(host, port)

		if err != nil {
		return err
		}

		i, _ := r.RowsAffected()

		if i != 1 {
		return errors.New("Should error rows")
		}


	return nil
}

// func MakeMigrations() error {
// 	db2 := GetConnection()
// 	defer db2.Close()

// 	q2 := `SELECT top 1 1 from domain`
// 	stmt2, err := db2.Prepare(q2)

// 	if err != nil {
// 		return err
// 	}
// 	defer stmt2.Close()

// 	r2, err := stmt2.Exec(q2)

// 	if err != nil {
// 		return err
// 	}

// 	i2, _ := r2.RowsAffected()
// 	fmt.Println("RowsAffected")
// 	fmt.Println(i2)

// 	// q := `CREATE TABLE IF NOT EXISTS notes (
// 	//         id INTEGER PRIMARY KEY AUTOINCREMENT,
//     //    		title VARCHAR(64) NULL,
//     //    		description VARCHAR(200) NULL
// 	//       );`

// 	return nil
// }


// func CreateNotesHandler(w http.ResponseWriter, r *http.Request) {
//     var note Note
	
//     err := json.NewDecoder(r.Body).Decode(&note)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }
//     err = note.Create()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
// }

// func UpdateNotesHandler(w http.ResponseWriter, r *http.Request) {
//     var note Note
// err := json.NewDecoder(r.Body).Decode(&note)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }
//     err = note.Update()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
// }

// func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
//     idStr := r.URL.Query().Get("id")
//     if idStr == "" {
//          http.Error(w, "Query id es requerido",
//              http.StatusBadRequest)
//          return
//     }
//     id, err := strconv.Atoi(idStr)
//     if err != nil {
//          http.Error(w, "Query id debe ser un número",
//              http.StatusBadRequest)
//          return
//     }
//     var note Note
//     err = note.Delete(id)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
// }

// func NotesHandler(w http.ResponseWriter, r *http.Request) {
//     switch r.Method {
//         case http.MethodPost:
//             CreateNotesHandler(w, r)
//         case http.MethodPut:
//             UpdateNotesHandler(w, r)
//         case http.MethodDelete:
//             DeleteNotesHandler(w, r)
//         default:
//             http.Error(w, "Metodo no permitido",
//                 http.StatusBadRequest)
//             return
//     }
// }

// type Note struct {
// 	ID          int       `json:"id,omitempty"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	CreatedAt   time.Time `json:"created_at,omitempty"`
// 	UpdatedAt   time.Time `json:"updated_at,omitempty"`
// }

// func (n Note) Create() error {
// 	db := GetConnection()

// 	q := `INSERT INTO notes (id,title, description)
// 			VALUES($1, $2, $3)`

// 	stmt, err := db.Prepare(q)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	r, err := stmt.Exec(n.ID,n.Title, n.Description)
// 	if err != nil {
// 		return err
// 	}
// 	if i, err := r.RowsAffected(); err != nil || i != 1 {
// 		return errors.New("ERROR: Se esperaba una fila afectada")
// 	}
// 	return nil
// }


// func (n *Note) GetByID(id int) (Note, error) {
// 	db := GetConnection()
// 	q := `SELECT
// 		id, title, description
// 		FROM notes WHERE id=$1`

// 	err := db.QueryRow(q, id).Scan(
// 		&n.ID, &n.Title, &n.Description,
// 	)
// 	if err != nil {
// 		return Note{}, err
// 	}

// 	return *n, nil
// }

// func (n Note) Update() error {
// 	db := GetConnection()
// 	q := `UPDATE notes set title=$1, description=$2
// 		WHERE id=$3`
// 	stmt, err := db.Prepare(q)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	r, err := stmt.Exec(n.Title, n.Description, n.ID)
// 	if err != nil {
// 		return err
// 	}
// 	if i, err := r.RowsAffected(); err != nil || i != 1 {
// 		return errors.New("ERROR: Se esperaba una fila afectada")
// 	}
// 	return nil
// }

// func (n Note) Delete(id int) error {
// 	db := GetConnection()

// 	q := `DELETE FROM notes
// 		WHERE id=$1`
// 	stmt, err := db.Prepare(q)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	r, err := stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}
// 	if i, err := r.RowsAffected(); err != nil || i != 1 {
// 		return errors.New("ERROR: Se esperaba una fila afectada")
// 	}
// 	return nil
// }