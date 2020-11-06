package persona

import (
	"context"
	"encoding/json"
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
	r.Method(http.MethodGet, "/{id}", getUserByHandler)

	getProductHandler := kithttp.NewServer(
		makeGetPersonsEndPoint(s),
		getPersonsRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	r.Method(http.MethodPost, "/paginated", getProductHandler)
	return r
}

//Decodificadores
func getPersonByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getPersonByIdRequest{
		PersonaId: chi.URLParam(r, "id"),
	}, nil
}

//Decodificador de los parametros del cuerpo
func getPersonsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getPersonsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request) //EL REQUEST QUE QUEREMOS DECODIFICAR ESTA EN BADY
	if err != nil {
		panic(err)
	}
	return request, nil
}
