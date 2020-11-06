package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func InitDB() *sql.DB {
	//connectionString := "root:hector@tcp(localhost:3306)/northwind"
	connectionString := "PI_IV/HECTOR@localhost:1521/ORCL"
	databaseConnection, err := sql.Open("oci8", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
