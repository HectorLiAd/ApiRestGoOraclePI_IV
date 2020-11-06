//Logica de negocio para consumir del repository
package persona

type Service interface {
	GetPersonById(param *getPersonByIdRequest) (*Persona, error)
}

type service struct {
	repo Repository
}

func NerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetPersonById(param *getPersonByIdRequest) (*Persona, error) {
	//Logica del negocio
	return s.repo.GetPersonById(param.PersonaId)
}
