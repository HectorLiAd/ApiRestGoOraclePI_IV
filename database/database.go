package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func InitDB() *sql.DB {
	//connectionString := "root:hector@tcp(localhost:3306)/northwind"
	connectionString := "DBII/123456@52.205.132.127:1521/orcl"
	databaseConnection, err := sql.Open("oci8", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
