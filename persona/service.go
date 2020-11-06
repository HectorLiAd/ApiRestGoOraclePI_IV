//Logica de negocio para consumir del repository
package persona

type Service interface {
	GetPersonById(param *getPersonByIdRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) (*PersonList, error)
}

type service struct {
	repo Repository
}

func NerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetPersonById(param *getPersonByIdRequest) (*Person, error) {
	//Logica del negocio
	return s.repo.GetPersonById(param.PersonaId)
}

func (s *service) GetPersons(params *getPersonsRequest) (*PersonList, error) {
	person, err := s.repo.GetPersons(params)
	if err != nil {
		panic(err)
	}
	totalPersons, err := s.repo.GetTotalPersons()
	if err != nil {
		panic(err)
	}
	return &PersonList{
		Data:         person,
		TotalRecords: totalPersons,
	}, nil
}
