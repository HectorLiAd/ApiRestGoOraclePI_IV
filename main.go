package main

import (
	_ "database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-oci8"

	"github.com/ApiRestGoOraclePI_IV/database"
	"github.com/ApiRestGoOraclePI_IV/usuario"
)

func main() {
	db := database.InitDB()
	//INSERTAR REGISTROS
	const sql = "begin PKG_CRUD_USU.SP_INGRESAR_USUARIO(:1, :2, :3, :4, :5, :6, :7, :8); end;"

	var nombre string = "Lewis"
	var nombre_personal string = "Lewis"
	var apellido_paterno string = "Orlan"
	var apellido_materno string = "Espinoza"
	var celular string = "95042142"
	var email string = "Lewis@gmail.com"
	var pwd string = "dfs415V1D5v5Vb"

	if _, err := db.Exec(sql, nil, nombre, nombre_personal, apellido_paterno, apellido_materno, celular, email, pwd); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	var usuarioRepository = usuario.NewRepository(db)
	var usuarioServicio = usuario.NerService(usuarioRepository)

	r := chi.NewRouter()
	r.Mount("/usuario", usuario.MakeHttpHandler(usuarioServicio))
	http.ListenAndServe(":3000", r)
}
