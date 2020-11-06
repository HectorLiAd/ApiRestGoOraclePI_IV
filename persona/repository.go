package persona

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetPersonById(personaId string) (*Persona, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}

func (repo *repository) GetPersonById(personaId string) (*Persona, error) {
	const sql = `SELECT * FROM PERSONA WHERE PERSONAS_ID = :1`
	row := repo.db.QueryRow(sql, personaId)
	persona := &Persona{}
	err := row.Scan(
		&persona.Id,
		&persona.Nombre,
		&persona.Apellido_paterno,
		&persona.Apellido_materno,
		&persona.Genero,
		&persona.Dni,
		&persona.Fecha_nacimiento,
		&persona.Edad,
	)
	if err != nil {
		fmt.Println("No hay resultados")
		panic(err)
	}
	return persona, err
}
