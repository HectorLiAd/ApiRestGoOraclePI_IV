package main

import (
	"fmt"

	"github.com/ApiRestGoOraclePI_IV/database"
)

func main() {
	databaseConnection := database.InitDB()

	var user string
	err := databaseConnection.QueryRow("SELECT PATERNO FROM CLIENTE").Scan(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successful 'as sysdba' connection. Current user is: %v\n", user)
}
