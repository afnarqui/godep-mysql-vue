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
	"github.com/satori/go.uuid"
	 "os"
	"path/filepath"
	"strings"
	"github.com/go-chi/chi"
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
