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
	 "os"
	"path/filepath"
	"strings"
	"github.com/go-chi/chi"
	"errors"
)
var Host string
var db *sql.DB
var domainnew = Domain{}
var domainold = Domain{}
var responsedatasearch = []Domain{}
var responsedatasearchcomparar = []Domaincomparar{}

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	 var err error

	 db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	     if err != nil {
         log.Fatal("error connecting to the database: ", err)
     }
	 
 return db
}

func (n *Domain) GetAllDomain() ([]Domain, error) {
	db := GetConnection()
	Host = "'"+Host+"'"

	q := "select distinct host,port,protocol,ispublic,status from domain where host="+string(Host)
	rows, err := db.Query(q)
	if err != nil {
		return []Domain{}, err
	}
	//,&bk.StartTime,&bk.TestTime,&bk.EngineVersion,&bk.CriteriaVersion
	defer rows.Close()
	bks := make([]Domain, 0)
	for rows.Next() {
		bk := Domain{}
		err := rows.Scan(&bk.Host, &bk.Port,&bk.Protocol,&bk.IsPublic,&bk.Status) 
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	return bks, nil 
}

func (n *Domain) GetDomain() ([]Domain, error) {
	db := GetConnection()
	
	q := "SELECT host,port FROM domain"
	rows, err := db.Query(q)
	if err != nil {
		return []Domain{}, err
	}
	defer rows.Close()
	bks := make([]Domain, 0)
	for rows.Next() {
		bk := Domain{}
		err := rows.Scan(&bk.Host, &bk.Port) 
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	return bks, nil 
}

func (n *Domaincomparar) GetDomaincomparar() ([]Domaincomparar, error) {
	db := GetConnection()
	
	q := "select domain.host as host,domainold.host as hostold,domain.port as port,domainold.port as portold FROM domain inner join domainold on domainold.host=domain.host"
	
	rows, err := db.Query(q)
	if err != nil {
		return []Domaincomparar{}, err
	}
	defer rows.Close()
	bks := make([]Domaincomparar, 0)
	for rows.Next() {
		bk := Domaincomparar{}
		err := rows.Scan(&bk.Host,&bk.Hostold ,&bk.Port,&bk.Portold) 
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	return bks, nil 
}

type Domain struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	Endpoints       []Endpoints `json:"endpoints"`
}

type Domaincomparar struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	Endpoints       []Endpoints `json:"endpoints"`
	Hostold            string      `json:"hostold"`
	Portold            int         `json:"portold"`
	Protocolold        string      `json:"protocolold"`
	IsPublicold        bool        `json:"isPublicold"`
	Statusold          string      `json:"statusold"`
	Endpointsold       []Endpoints `json:"endpointsold"`
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

	MakeMigrations()
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	
	log.Println("Corriendo en http://localhost:8081")
	r := chi.NewRouter()

	r.Get("/buscardomaincomparar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("entro en buscardomaincomparar")

		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		n := new(Domaincomparar)
		domain, err := n.GetDomaincomparar()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		dataDomain, err := json.Marshal(domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		responsedata:= []Domaincomparar{}
	
		errrs := json.Unmarshal([]byte(dataDomain), &responsedata)
		if errrs != nil {
			fmt.Println(errrs)
		}
		fmt.Println(responsedata)
		if len(responsedata) > 0 {
			responsedatasearchcomparar = responsedata
		} else {
		}
		json.NewEncoder(w).Encode(responsedatasearchcomparar)
})
	r.Get("/buscardomain", func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		n := new(Domain)
		domain, err := n.GetDomain()
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
		if len(responsedata) > 0 {
			fmt.Println("print")
			responsedatasearch = responsedata
		} else {
		}
		json.NewEncoder(w).Encode(responsedatasearch)
})
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
			 fmt.Println(i,v)
			 data.Host = v.Host
			 Host = v.Host
			 data.Port = v.Port
			 data.Protocol = v.Protocol
			 data.IsPublic = v.IsPublic
			 data.Status = v.Status
			//  data.StartTime = v.StartTime
			//  data.TestTime = v.TestTime
			//  data.EngineVersion = v.EngineVersion
			//  data.CriteriaVersion = v.CriteriaVersion
			
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
		domainnew = data
	
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
			var dataupdate Domain
			domainold = responsedata[0]
			fmt.Println("debo actualizar eliminar de todo")
			err = dataupdate.DeleteDomain()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}		
		} else {
			var datanew Domain
			fmt.Println("debo crear")
			err = datanew.CreateDomain()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(domainnew)
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
	direccion := ":8081" 
	fmt.Println("Servidor listo escuchando en " + direccion)

	log.Fatal(http.ListenAndServe(direccion+"/public/index.html", nil))
	
}

func (n Domain) CreateDomain() error {

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
	var protocol = domainnew.Protocol 
	var isPublic = domainnew.IsPublic
	fmt.Println("IsPublic")
	fmt.Println(domainnew.IsPublic)
	var status = domainnew.Status
	// var startTime = domainnew.StartTime
	// var testTime = domainnew.TestTime
	// var engineVersion = domainnew.EngineVersion   
	// var criteriaVersion = domainnew.CriteriaVersion
	// var endpointss = domainnew.Endpoints

	// q := `INSERT INTO 
	// domain(host,port,protocol,isPublic,status)
	// VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	q := "INSERT INTO domain(host,port,protocol,ispublic,status) VALUES ($1,$2,$3,$4,$5)"
	// select host,port,protocol,ispublic,status,starttime,testtime,engineversion,criteriaversion from domain;
		db := GetConnection()
		defer db.Close()
		fmt.Println("va a guardar")
		fmt.Println(q)
		stmt, err := db.Prepare(q)

		if err != nil {
		return err
		}
		defer stmt.Close()
	     // select host,port,protocol,ispublic,status,starttime,testtime,engineversion,criteriaversion from domain;
		// r, err := stmt.Exec(host,port,protocol,isPublic,status,startTime,testTime,engineVersion,criteriaVersion,endpointss)

		r, err := stmt.Exec(host,port,protocol,isPublic,status)
		if err != nil {
		return err
		}

		i, _ := r.RowsAffected()

		if i != 1 {
		return errors.New("Should error rows")
		}


	return nil
}

func (n Domain) DeleteDomain() error {


	dbdomain := GetConnection()

	var host = domainnew.Host

	qdomain := `DELETE FROM domain
		WHERE host=$1`
	stmtdomain, errdomain := dbdomain.Prepare(qdomain)
	if errdomain != nil {
		// return errdomain
	}
	defer stmtdomain.Close()

	rdomain, errdomain := stmtdomain.Exec(host)
	if errdomain != nil {
		// return errdomain
	}
	if idomain, errdomain := rdomain.RowsAffected(); errdomain != nil || idomain != 1 {
		//return errors.New("ERROR: Se esperaba una fila afectada")
	}

	dbdomainold := GetConnection()

	qdomainold := `DELETE FROM domainold
		WHERE host=$1`
	stmtdomainold, errdomainold := dbdomainold.Prepare(qdomainold)
	if errdomainold != nil {
		//return errdomainold
	}
	defer stmtdomainold.Close()

	rdomainold, errdomainold := stmtdomainold.Exec(host)
	if errdomainold != nil {
		//return errdomainold
	}
	if idomainold, errdomainold := rdomainold.RowsAffected(); errdomainold != nil || idomainold != 1 {
		//return errors.New("ERROR: Se esperaba una fila afectada")
	}

	var portnewdomain = domainnew.Port
	var protocolnewdomain = domainnew.Protocol 
	var isPublicnewdomain = domainnew.IsPublic 
	var statusnewdomain = domainnew.Status
	// var starttimenewdomain = domainnew.StartTime
	// var testtimenewdomain = domainnew.TestTime
	// var engineversionnewdomain = domainnew.EngineVersion
	// var criteriaversionnewdomain = domainnew.CriteriaVersion
	// ,starttime,testtime,engineversion,criteriaversion
	// ,$6,$7,$8,$9
	qnewdomain := `INSERT INTO 
	domain(host,port,protocol,ispublic,status)
	VALUES ($1,$2,$3,$4,$5)`
     
		dbnewdomain := GetConnection()
		defer dbnewdomain.Close()

		stmtnewdomain, errnewdomain := dbnewdomain.Prepare(qnewdomain)

		if errnewdomain != nil {
		return errnewdomain
		}
		defer stmtnewdomain.Close()
		
		
		//,starttimenewdomain,testtimenewdomain,engineversionnewdomain,criteriaversionnewdomain
		rnewdomain, errnewdomain := stmtnewdomain.Exec(host,portnewdomain,protocolnewdomain,isPublicnewdomain,statusnewdomain)
		if errnewdomain != nil {
		return errnewdomain
		}

		inewdomain, _ := rnewdomain.RowsAffected()

		if inewdomain != 1 {
		return errors.New("Should error rows newdomain")
		}


		var portolddomain = domainold.Port
		var protocololddomain = domainold.Protocol 
		var isPublicolddomain = domainold.IsPublic 
		var statusolddomain = domainold.Status
		// var starttimeolddomain = domainold.StartTime
		// var testtimeolddomain = domainold.TestTime
		// var engineversionolddomain = domainold.EngineVersion
		// var criteriaversionolddomain = domainold.CriteriaVersion
		//,starttime,testtime,engineversion,criteriaversion
		// ,$6,$7,$8,$9
		qolddomain := `INSERT INTO 
		domainold(host,port,protocol,ispublic,status)
		VALUES ($1,$2,$3,$4,$5)`

			dbolddomain := GetConnection()
			defer dbolddomain.Close()
	
			stmtolddomain, errolddomain := dbolddomain.Prepare(qolddomain)
	
			if errolddomain != nil {
			return errolddomain
			}
			defer stmtolddomain.Close()
			//,starttimeolddomain,testtimeolddomain,engineversionolddomain,criteriaversionolddomain
			rolddomain, errolddomain := stmtolddomain.Exec(host,portolddomain,protocololddomain,isPublicolddomain,statusolddomain)
			if errolddomain != nil {
			return errolddomain
			}
	
			iolddomain, _ := rolddomain.RowsAffected()
	
			if iolddomain != 1 {
			return errors.New("Should error rows olddomain")
			}		

			//,starttime,testtime,engineversion,criteriaversion
			// ,$6,$7,$8,$9
			qhistorydomain := `INSERT INTO 
			domainhistory(host,port,protocol,ispublic,status)
							VALUES ($1,$2,$3,$4,$5)`
		
				dbhistorydomain := GetConnection()
				defer dbhistorydomain.Close()
		
				stmthistorydomain, errhistorydomain := dbhistorydomain.Prepare(qhistorydomain)
		
				if errhistorydomain != nil {
				return errhistorydomain
				}
				defer stmthistorydomain.Close()
				//,starttimeolddomain,testtimeolddomain,engineversionolddomain,criteriaversionolddomain
				rhistorydomain, errhistorydomain := stmthistorydomain.Exec(host,portolddomain,protocololddomain,isPublicolddomain,statusolddomain)
				if errhistorydomain != nil {
				return errhistorydomain
				}
		
				ihistorydomain, _ := rhistorydomain.RowsAffected()
		
				if ihistorydomain != 1 {
				return errors.New("Should error rows historydomain")
				}		
	

	return nil
}

 func MakeMigrations() error {
	db2 := GetConnection()
	defer db2.Close()

	q2 := `SELECT host,port from domain`
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
	if (i2 > 0){
		fmt.Println("no hace nada")
		return nil
	}
	fmt.Println("va a correr la migración")

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

		if _, err := db.Exec(
			`DROP TABLE IF EXISTS domain`); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(
			`DROP TABLE IF EXISTS domainold`); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(
			`DROP TABLE IF EXISTS domainhistory`); err != nil {
			log.Fatal(err)
		}
		// StartTime       VARCHAR(120) NULL,
		// TestTime        TIMESTAMP DEFAULT now(),
		// EngineVersion   VARCHAR(120)  NULL,
		// CriteriaVersion VARCHAR(120)  NULL,
		// Endpoints       VARCHAR(8000) NULL  
		
		if _, err := db.Exec(
			`CREATE TABLE IF NOT EXISTS domain (
					Host VARCHAR(120) NULL,
					Port INT NULL,
					Protocol VARCHAR(120) NULL,
					IsPublic BOOL NULL,
					Status   VARCHAR(80) NULL,
					Endpoints       VARCHAR(8000) NULL 
					)`); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(
			`CREATE TABLE IF NOT EXISTS domainold (
				Host VARCHAR(120) NULL,
				Port INT NULL,
				Protocol VARCHAR(120) NULL,
				IsPublic BOOL NULL,
				Status   VARCHAR(80) NULL,
				Endpoints       VARCHAR(8000) NULL 
				)`); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(
			`CREATE TABLE IF NOT EXISTS domainhistory (
				Host VARCHAR(120) NULL,
				Port INT NULL,
				Protocol VARCHAR(120) NULL,
				IsPublic BOOL NULL,
				Status   VARCHAR(80) NULL,
				Endpoints       VARCHAR(8000) NULL 
				)`); err != nil {
			log.Fatal(err)
		}



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
		// 			endpoints
		// 		) VALUES (
		// 			'www.google.com',
		// 			444,
		// 			'http',
		// 			false,
		// 			'READY edit',
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
		// 			  ]}')`); err != nil {
		// 	log.Fatal(err)
		// }
		 
	return nil
 }
