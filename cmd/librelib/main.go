package main

import (
	"fmt"

	"github.com/DenHax/LibreLib-back/internal/config"
	"github.com/DenHax/LibreLib-back/internal/database"
	"github.com/DenHax/LibreLib-back/internal/router"
)

func main() {
	fmt.Println("LibreLib Backend")

	cfg := config.MustConfig()

	db := database.Connect()
	defer db.Close()

	//TODO: Config init
	//TODO: Logger init
	//TOD: PSQL init
	//TODO: Server init

	r := router.GetRouter()

	// http.HandleFunc("/set-name", setNameHandler)
	// http.HandleFunc("/get-name", getNameHandler)
	http.ListenAndServe(":8080", r)
}
