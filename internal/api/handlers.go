package api

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
)

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"name": "hello"})
}
