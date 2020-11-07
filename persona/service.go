//Logica de negocio para consumir del repository
package persona

type Service interface {
	GetPersonById(param *getPersonByIdRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) (*PersonList, error)
	InsertPerson(params *getAddPersonRequest) (int64, error)
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
	//Logica del negocio
	params.Limit = params.Offset + params.Limit
	params.Offset = params.Offset + 1

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

func (s *service) InsertPerson(params *getAddPersonRequest) (int64, error) {
	return s.repo.InsertPerson(params)
}
