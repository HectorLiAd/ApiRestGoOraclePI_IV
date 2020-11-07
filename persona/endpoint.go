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
	nombre           string
	apellido_paterno string
	apellido_materno string
	genero           string
	dni              string
	fecha_nacimiento string
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

		result, err := s.InsertPerson(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}

	return addPersonEndpoint
}
