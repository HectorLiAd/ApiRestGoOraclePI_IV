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

	getPersonByHandler := kithttp.NewServer(
		makeGetPersonByIdEndPoint(s),
		getPersonByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getPersonByHandler)

	getPersonHandler := kithttp.NewServer(
		makeGetPersonsEndPoint(s),
		getPersonsRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getPersonHandler)

	addPersonHandler := kithttp.NewServer(
		makeAddPersonEndpoint(s),
		addPersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addPersonHandler)

	//Actualizar personas
	updatePersonHandler := kithttp.NewServer(
		makeUpdatePersonEndpoint(s),
		updatePersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/", updatePersonHandler)

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

func addPersonRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddPersonRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

//Decodificar los PROPIEDADES QUE ESTAN EN EL BY DEL REQUEST
func updatePersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updatePersonRequest{}
	//LAS PROPIEDADES DEL BODY REQUEST SE MAPEAN PARA PODER OBTENER EL FORMATO DE NUESTRA ESTRUCTURA INDICADA
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}
