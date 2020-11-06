package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getPersonByIdRequest struct {
	PersonaId string
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
