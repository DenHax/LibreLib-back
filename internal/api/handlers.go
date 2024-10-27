package api

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
)

func GetPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Name: "u name"}
	json.NewEncoder(w).Encode(response)
}
