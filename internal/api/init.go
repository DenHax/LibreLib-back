package api

import (
	"encoding/json"
	"fmt"
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

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"name": "hello"})
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main page")
}

func (h Handler) Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	r.HandleFunc("/info", GetNameHandler)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.HeaderMiddleware)
	api = h.NewApi(api)
	return r
}
