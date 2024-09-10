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
	var data struct {
		Name string `json:"name"`
	}
	if r.Method == http.MethodPost {
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
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
	db := database.Connect()
	defer db.Close()
	rows := database.Select_data(db, `SELECT 
    b.name,
    b.author,
    b.genre,
    b.year,
    p.price,
	p.id,
	p.salesmanid,
	p.type
FROM 
    "cart-product" cp
JOIN 
    "product" p ON cp.productid = p.id
JOIN 
    "book" b ON p.id = b.id
JOIN 
    "customer" c ON cp.cartid = c.id
WHERE 
    c.id = 3 
AND 
    p.type = 'book';`)
	defer rows.Close()
	books := database.Set_Book_data(rows)
	for _, book := range books {
		fmt.Println(book.Name, book.Author, book.Genre, book.Year, book.Price, book.Product.ID, book.Product.SalesmanID, book.Product.Type)
	}
	JSON_str := database.ToJSON(books)
	fmt.Println(JSON_str)

	// http.HandleFunc("/set-name", setNameHandler)
	// http.HandleFunc("/get-name", getNameHandler)
	// http.ListenAndServe(":8080", nil)
}
