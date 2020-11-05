package usuario

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getUsuarioByHandler := kithttp.NewServer(
		makeGetUsuarioByIdEndPoint(s),
		getUsuarioByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(
		http.MethodGet,
		"/{id}",
		getUsuarioByHandler,
	)

	return r
}

func getUsuarioByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getUsuarioByIdRequest{
		UsuarioId: chi.URLParam(r, "id"),
	}, nil
}
