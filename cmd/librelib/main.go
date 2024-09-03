package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/DenHax/LibreLib-back/internal/database"
)

var (
	name  string
	mutex sync.Mutex
)

func setNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var data struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mutex.Lock()
		name = data.Name
		mutex.Unlock()
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"name": name})
}

func main() {
	fmt.Println("LibreLib Backend")
	database.Start()
	// http.HandleFunc("/set-name", setNameHandler)
	// http.HandleFunc("/get-name", getNameHandler)

	// http.ListenAndServe(":8080", nil)
}
