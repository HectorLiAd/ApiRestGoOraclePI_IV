package usuario

type Usuario struct {
	Usuario_id               string `json:"usuario_id"`
	Usuario_nombre           string `json:"Usuario_nombre"`
	Usuario_nombre_personal  string `json:"Usuario_nombre_personal"`
	Usuario_apellido_paterno string `json:"Usuario_apellido_paterno"`
	Usuario_apellido_materno string `json:"Usuario_apellido_materno"`
	Usuario_celular          string `json:"Usuario_celular"`
	Usuario_email            string `json:"Usuario_email"`
	Usuario_contrasenia      string `json:"Usuario_contrasenia"`
}
