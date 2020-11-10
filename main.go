package main

import (
	_ "database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-oci8"

	"github.com/ApiRestGoOraclePI_IV/database"
	"github.com/ApiRestGoOraclePI_IV/helper"
	"github.com/ApiRestGoOraclePI_IV/persona"
)

func main() {
	db := database.InitDB()

	defer db.Close()
	var personaRepository = persona.NewRepository(db)
	var personaServicio = persona.NerService(personaRepository)

	r := chi.NewRouter()

	r.Use(helper.GetCors().Handler)
	r.Mount("/persona", persona.MakeHttpHandler(personaServicio))
	http.ListenAndServe(":3000", r)

}
