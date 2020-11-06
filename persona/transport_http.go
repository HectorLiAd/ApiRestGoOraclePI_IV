package persona

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getUserByHandler := kithttp.NewServer(
		makeGetPersonByIdEndPoint(s),
		getPersonByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(
		http.MethodGet,
		"/{id}",
		getUserByHandler,
	)

	return r
}

func getPersonByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getPersonByIdRequest{
		PersonaId: chi.URLParam(r, "id"),
	}, nil
}
