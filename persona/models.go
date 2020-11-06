package persona

import "time"

type Person struct {
	Id               string    `json:"id"`
	Nombre           string    `json:"nombre_personal"`
	Apellido_paterno string    `json:"apellido_paterno"`
	Apellido_materno string    `json:"apellido_materno"`
	Genero           string    `json:"Genero"`
	Dni              string    `json:"dni"`
	Fecha_nacimiento time.Time `json:"fecha_nacimiento"`
	Edad             int       `json:"edad"`
}

type PersonList struct {
	Data         []*Person `json:"data"`
	TotalRecords int       `json:"totalRecords"`
}
