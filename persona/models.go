package persona

type Person struct {
	Id               string `json:"id"`
	Nombre           string `json:"nombre_personal"`
	Apellido_paterno string `json:"apellido_paterno"`
	Apellido_materno string `json:"apellido_materno"`
	Genero           string `json:"Genero"`
	Dni              string `json:"dni"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
	Estado           int    `json:"estado"`
}

type PersonList struct {
	Data         []*Person `json:"data"`
	TotalRecords int       `json:"totalRecords"`
}

type StatusPerson struct {
	Error  string `json:"Error"`
	Estado string `json:"Status"`
}
