package main

import (
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	service Service
}

func NewApiServer(service Service) *ApiServer {
	return &ApiServer{
		service: service,
	}
}

func (apiServer *ApiServer) Start(listenAddress string) error {
	http.HandleFunc("/status", apiServer.handleStatus)
	http.HandleFunc("/chuckNorrisJoke", apiServer.handleGetChuckNorrisJoke)

	return http.ListenAndServe(listenAddress, nil)
}

func (apiServer *ApiServer) handleStatus(respWriter http.ResponseWriter, _ *http.Request) {
	respWriter.Write([]byte("OK"))
}

func (apiServer *ApiServer) handleGetChuckNorrisJoke(respWriter http.ResponseWriter, req *http.Request) {
	chuckNorrisJoke, err := apiServer.service.GetChuckNorrisJoke(req.Context())
	if err != nil {
		http.Error(respWriter, "Failed to fetch Chuck Norris joke", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(respWriter).Encode(chuckNorrisJoke)
}
