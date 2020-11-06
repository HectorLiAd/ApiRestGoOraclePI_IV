package main

import (
	_ "database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-oci8"

	"github.com/ApiRestGoOraclePI_IV/database"
	"github.com/ApiRestGoOraclePI_IV/persona"
)

func main() {
	db := database.InitDB()

	defer db.Close()
	var personaRepository = persona.NewRepository(db)
	var personaServicio = persona.NerService(personaRepository)

	r := chi.NewRouter()
	r.Mount("/persona", persona.MakeHttpHandler(personaServicio))
	http.ListenAndServe(":3000", r)

}

func pruebasBD() {
	// INSERTAR REGISTROS
	// codigo := "2020PER000013"
	// persona := &persona.Persona{}

	// const sql = `SELECT * FROM PERSONA WHERE PERSONAS_ID = :1`
	// row := db.QueryRow(sql, codigo)

	// err := row.Scan(
	// 	&persona.Id,
	// 	&persona.Nombre,
	// 	&persona.Apellido_paterno,
	// 	&persona.Apellido_materno,
	// 	&persona.Genero,
	// 	&persona.Dni,
	// 	&persona.Fecha_nacimiento,
	// 	&persona.Edad,
	// )

	// if err != nil {
	// 	fmt.Println("Error running query")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf(persona.Id)

	// const sql = "INSERT INTO PERSONA ( NOMBRE, APELLIDO_P, APELLIDO_M, ESTADO ) VALUES('02', :1, :2, :3, :4);"

	// if _, err := db.Exec(sql, "s", "efs", "bd", 0); err != nil {
	// 	log.Fatal(err)
	// }
}
