package delivery

import (
	"github.com/gorilla/mux"
	"backend/api"
)

func RegisterHTTPEndpoint(router *mux.Router, useCase api.UseCase) {
	h := newHandler(useCase)

	router.HandleFunc("/api/getObjects", h.getObjects)
}