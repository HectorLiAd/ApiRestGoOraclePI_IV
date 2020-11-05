//Logica de negocio para consumir del repository
package usuario

type Service interface {
	GetUsuarioById(param *getUsuarioByIdRequest) (*Usuario, error)
}

type service struct {
	repo Repository
}

func NerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUsuarioById(param *getUsuarioByIdRequest) (*Usuario, error) {
	//Logica del negocio
	return s.repo.GetUsuarioById(param.UsuarioId)
}
