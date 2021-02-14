package delivery

import (
	"backend/api"
	"net/http"
	"encoding/json"
	"strconv"
)

type handler struct {
	useCase api.UseCase
}

func newHandler(useCase api.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) getObjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")

	firstObjectNumber, err := strconv.Atoi(r.URL.Query().Get("first"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	objectsCount, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	objects, err := h.useCase.GetObjects(firstObjectNumber, objectsCount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} 
	body, err := json.MarshalIndent(objects, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}
}