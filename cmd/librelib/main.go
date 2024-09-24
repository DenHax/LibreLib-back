package main

import (
	"fmt"
	"net/http"

	"github.com/DenHax/LibreLib-back/internal/config"
	"github.com/DenHax/LibreLib-back/internal/router"
)

func main() {
	fmt.Println("LibreLib Backend")

	cfg := config.MustConfig()
	// storage, err := postgres.New(cfg.Storage)
	// if err != nil {
	// 	fmt.Println("failed to init storage ",err)
	// 	os.Exit(1)
	// }
	// defer storage.Close()

	//TODO: Config init
	//TODO: Logger init
	//TOD: PSQL init
	//TODO: Server init

	r := router.GetRouter()

	// http.HandleFunc("/set-name", setNameHandler)
	// http.HandleFunc("/get-name", getNameHandler)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}
