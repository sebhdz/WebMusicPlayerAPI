package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Conectar() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/player?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al iniciar base de datos: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error: No se pudo conectar con la base de datos: ", err)
	}
	fmt.Println("Connected to database")
}
