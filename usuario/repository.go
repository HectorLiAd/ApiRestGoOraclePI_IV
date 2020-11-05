package usuario

import "database/sql"

type Repository interface {
	GetUsuarioById(usuarioId string) (*Usuario, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}

func (repo *repository) GetUsuarioById(usuarioId string) (*Usuario, error) {
	const sql = `SELECT * FROM USUARIO WHERE USUARIO_ID = :1`
	row := repo.db.QueryRow(sql, usuarioId)
	usuario := &Usuario{}
	err := row.Scan(
		&usuario.Usuario_id,
		&usuario.Usuario_nombre_personal,
		&usuario.Usuario_nombre,
		&usuario.Usuario_apellido_paterno,
		&usuario.Usuario_apellido_materno,
		&usuario.Usuario_celular,
		&usuario.Usuario_email,
		&usuario.Usuario_contrasenia,
	)
	if err != nil {
		panic(err)
	}
	return usuario, err
}
