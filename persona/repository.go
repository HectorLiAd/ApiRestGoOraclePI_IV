package persona

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetPersonById(personaId string) (*Person, error)
	GetPersons(params *getPersonsRequest) ([]*Person, error)
	GetTotalPersons() (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}

func (repo *repository) GetPersonById(personaId string) (*Person, error) {
	const sql = `SELECT * FROM PERSONA WHERE PERSONAS_ID = :1`
	row := repo.db.QueryRow(sql, personaId)
	persona := &Person{}
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

func (repo *repository) GetPersons(params *getPersonsRequest) ([]*Person, error) {
	const sql = `SELECT PERSONAS_ID,NOMBRE,APELLIDO_P,APELLIDO_M,GENERO,DNI,FECHA_NACIMIENTO,ESTADO FROM(
				SELECT ROWNUM COD, E.* FROM PERSONA E
				) WHERE COD BETWEEN :1 AND :2 
				ORDER BY COD`
	result, err := repo.db.Query(sql, (params.Offset + 1), (params.Offset + params.Limit))

	if err != nil {
		panic(err)
	}

	var persons []*Person
	for result.Next() {
		persona := &Person{}
		err := result.Scan(
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
			panic(err)
		}
		persons = append(persons, persona)
	}
	return persons, nil
}

func (repo *repository) GetTotalPersons() (int, error) {
	const sql = `SELECT COUNT(*) FROM PERSONA`
	var total int

	row := repo.db.QueryRow(sql)

	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}
