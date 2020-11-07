package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getPersonByIdRequest struct {
	PersonaId string
}

type getPersonsRequest struct {
	Limit  int //CUANTOS REGISTROS TRAER
	Offset int //DE QUE NUMERO DE FILA INICIARA LA CONSULTA
}

type getAddPersonRequest struct {
	Nombre           string
	Apellido_paterno string
	Apellido_materno string
	Genero           string
	Dni              string
	Fecha_nacimiento string
}

type updatePersonRequest struct {
	Id               string
	Nombre           string
	Apellido_paterno string
	Apellido_materno string
	Genero           string
	Dni              string
	Fecha_nacimiento string
}

func makeGetPersonByIdEndPoint(s Service) endpoint.Endpoint {
	getPersonById := func(ctx context.Context, request interface{}) (interface{}, error) {
		rep := request.(getPersonByIdRequest)
		persona, err := s.GetPersonById(&rep)
		if err != nil {
			panic(err)
		}
		return persona, nil
	}
	return getPersonById
}

func makeGetPersonsEndPoint(s Service) endpoint.Endpoint {
	getPersonsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPersonsRequest) //Conversion del request al tipo getPersonsRequest
		result, err := s.GetPersons(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getPersonsEndPoint
}

func makeAddPersonEndpoint(s Service) endpoint.Endpoint {
	addPersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddPersonRequest)
		persona_id, err := s.InsertPerson(&req)

		if err != nil {
			panic(err)
		}
		return persona_id, nil
	}

	return addPersonEndpoint
}

func makeUpdatePersonEndpoint(s Service) endpoint.Endpoint {
	updatePersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updatePersonRequest)
		r, err := s.UpdatePerson(&req)

		if err != nil {
			panic(err)
		}
		return r, nil
	}
	return updatePersonEndpoint
}
