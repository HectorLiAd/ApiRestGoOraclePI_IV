//Logica de negocio para consumir del repository
package persona

// "time"
// "fmt"

type Service interface {
	GetPersonById(param *getPersonByIdRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) (*PersonList, error)
	InsertPerson(params *getAddPersonRequest) (int64, error)
	UpdatePerson(params *updatePersonRequest) (int64, error)
	DeletePerson(param *deletePersonRequest) (int64, error)
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
	persona, err := s.repo.GetPersonById(param.PersonaId)
	return persona, err
}

func (s *service) GetPersons(params *getPersonsRequest) (*PersonList, error) {
	//Logica del negocio
	params.Limit = (params.Offset + params.Limit - 1)
	person, err := s.repo.GetPersons(params)
	if err != nil {
		return nil, err
	}
	totalPersons, err := s.repo.GetTotalPersons()
	if err != nil {
		return nil, err
	}
	return &PersonList{
		Data:         person,
		TotalRecords: totalPersons,
	}, err
}

func (s *service) InsertPerson(params *getAddPersonRequest) (int64, error) {
	//Usar sub-string segun
	return s.repo.InsertPerson(params)
}

func (s *service) UpdatePerson(params *updatePersonRequest) (int64, error) {
	return s.repo.UpdatePerson(params)
}

func (s *service) DeletePerson(param *deletePersonRequest) (int64, error) {
	return s.repo.DeletePerson(param)
}
