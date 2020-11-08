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

	//Obtener personas por su identificador
	getPersonByHandler := kithttp.NewServer(
		makeGetPersonByIdEndPoint(s),
		getPersonByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getPersonByHandler)

	//Obtener personas paginadas
	getPersonHandler := kithttp.NewServer(
		makeGetPersonsEndPoint(s),
		getPersonsRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getPersonHandler)

	//Agregar a una persona
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

	//Eliminar PERSONA
	deletePersonHandler := kithttp.NewServer(
		makeDeletePersonEndPoint(s),
		deletePersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodDelete, "/{id}", deletePersonHandler)
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

	return request, err
}

func addPersonRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddPersonRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

//Decodificar los PROPIEDADES QUE ESTAN EN EL BY DEL REQUEST
func updatePersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updatePersonRequest{}
	//LAS PROPIEDADES DEL BODY REQUEST SE MAPEAN PARA PODER OBTENER EL FORMATO DE NUESTRA ESTRUCTURA INDICADA
	err := json.NewDecoder(r.Body).Decode(&request)

	return request, err
}

func deletePersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deletePersonRequest{
		PersonaId: chi.URLParam(r, "id"),
	}, nil
}
