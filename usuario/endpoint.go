package usuario

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getUsuarioByIdRequest struct {
	UsuarioId string
}

func makeGetUsuarioByIdEndPoint(s Service) endpoint.Endpoint {
	getUsuarioById := func(ctx context.Context, request interface{}) (interface{}, error) {
		rep := request.(getUsuarioByIdRequest)
		usuario, err := s.GetUsuarioById(&rep)
		if err != nil {
			panic(err)
		}
		return usuario, nil
	}
	return getUsuarioById
}
