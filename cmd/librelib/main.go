package main

import (
	"fmt"

	"github.com/DenHax/LibreLib-back/internal/config"
	"github.com/DenHax/LibreLib-back/internal/database"
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

}
