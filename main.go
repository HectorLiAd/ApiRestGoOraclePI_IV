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

/*
type addPersonRequest struct {
	Id               string
	Nombre           string
	Apellido_paterno string
	Apellido_materno string
	Genero           string
	Dni              string
}

func pruebasBD(db *sql.DB) {

	// INSERTAR REGISTROS
	persona := &addPersonRequest{}
	persona.Id = "ffs"
	persona.Nombre = "ffs"
	persona.Apellido_paterno = "ffs"
	persona.Apellido_materno = "ffs"
	persona.Genero = "m"
	persona.Dni = "ffs"
	dato := 0

	var salida *sql.Out
	salida.Dest = &dato

	const sql = `DECLARE
					ST_PERSONA PERSONA%ROWTYPE;
				BEGIN
					ST_PERSONA.NOMBRE := :1;
					ST_PERSONA.APELLIDO_P := :2;
					ST_PERSONA.APELLIDO_M := :3 ;
					ST_PERSONA.GENERO := :4;
					ST_PERSONA.DNI := :5;
					PKG_CRUD_PERSONA.SPU_AGREGAR_PERSONA(ST_PERSONA, :6);
				END;`
	_, err := db.Exec(sql, &persona.Nombre, &persona.Apellido_paterno, &persona.Apellido_materno,
		&persona.Genero, &persona.Dni, salida)
	if err != nil {
		panic(err)
	}

	fmt.Println("Salida: ", dato)
}
*/
