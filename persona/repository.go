package persona

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetPersonById(personaId string) (*Person, error)
	GetPersons(params *getPersonsRequest) ([]*Person, error)
	GetTotalPersons() (int, error)
	InsertPerson(params *getAddPersonRequest) (int64, error)
	UpdatePerson(params *updatePersonRequest) (int64, error)
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
	const sql = `SELECT * FROM PERSONA WHERE PERSONA_ID = :1`
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
		&persona.Estado,
	)
	if err != nil {
		fmt.Println("No hay resultados")
		panic(err)
	}
	return persona, err
}

func (repo *repository) GetPersons(params *getPersonsRequest) ([]*Person, error) {
	const sql = `SELECT PERSONA_ID,NOMBRE,APELLIDO_P,APELLIDO_M,GENERO,DNI,FECHA_NACIMIENTO,ESTADO FROM(
				SELECT ROWNUM COD, E.* FROM PERSONA E
				) WHERE COD BETWEEN :1 AND :2 
				ORDER BY COD`
	result, err := repo.db.Query(sql, params.Offset, params.Limit)

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
			&persona.Estado,
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

func (repo *repository) InsertPerson(params *getAddPersonRequest) (int64, error) {
	var status_query int
	const sql = `DECLARE
					ST_PERSONA PERSONA%ROWTYPE;
				BEGIN
					ST_PERSONA.NOMBRE := :1;
					ST_PERSONA.APELLIDO_P := :2;
					ST_PERSONA.APELLIDO_M := :3 ;
					ST_PERSONA.GENERO := :4;
					ST_PERSONA.DNI := :5;
					ST_PERSONA.FECHA_NACIMIENTO := :6;
					PKG_CRUD_PERSONA.SPU_AGREGAR_PERSONA(ST_PERSONA, :7);
				END;`
	result, err := repo.db.Exec(sql, params.Nombre, params.Apellido_paterno,
		params.Apellido_materno, params.Genero, params.Dni,
		params.Fecha_nacimiento, status_query)

	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdatePerson(params *updatePersonRequest) (int64, error) {
	var status_query int
	const sql = `DECLARE
					ST_PERSONA PERSONA%ROWTYPE;
				BEGIN
					ST_PERSONA.PERSONA_ID := :1;
					ST_PERSONA.NOMBRE := :2;
					ST_PERSONA.APELLIDO_P := :3;
					ST_PERSONA.APELLIDO_M := :4 ;
					ST_PERSONA.GENERO := :5;
					ST_PERSONA.DNI := :6;
					ST_PERSONA.FECHA_NACIMIENTO := :7;
					PKG_CRUD_PERSONA.SPU_ACTUALIZAR_PERSONA(ST_PERSONA, :8);
				END;`
	result, err := repo.db.Exec(sql, params.Id, params.Nombre, params.Apellido_paterno,
		params.Apellido_materno, params.Genero, params.Dni,
		params.Fecha_nacimiento, status_query)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}
