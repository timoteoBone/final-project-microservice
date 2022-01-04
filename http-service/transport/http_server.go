package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
	"github.com/timoteoBone/final-project-microservice/http-service/endpoints"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPSrv(endpoint endpoints.Endpoints) http.Handler {
	rt := mux.NewRouter()

	rt.Methods("POST").Path("/api/createUs").Handler(httptransport.NewServer(
		endpoint.CreateUs,
		decodeCreateUserReq,
		encodeCreateUserResp,
	))

	rt.Methods("GET").Path("/api/getUs/{id}").Handler(httptransport.NewServer(
		endpoint.GetUs,
		decodeGetUserReq,
		encodeGetUserResp,
	))

	return rt
}

func decodeCreateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {

	var request entities.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCreateUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(wr).Encode(response)
}

func decodeGetUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entities.GetUserRequest
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("unvalid request type")
	}

	request.UserID = id
	return request, nil
}

func encodeGetUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(wr).Encode(response)
}
