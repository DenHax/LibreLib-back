package api

import (
	"encoding/json"
	"net/http"

	"github.com/DenHax/LibreLib-back/internal/middleware"
	"github.com/DenHax/LibreLib-back/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

type Response struct {
	Name string `json:"name"`
}

func GetPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Name: "u name"}
	json.NewEncoder(w).Encode(response)
}

func (h Handler) Init() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.HeaderMiddleware)
	api = h.NewApi(api)
	return r
}
