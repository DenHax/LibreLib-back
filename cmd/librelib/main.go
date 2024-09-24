package main

import (
	"fmt"
	"net/http"

	"github.com/DenHax/LibreLib-back/internal/api/v1/routers"
	"github.com/DenHax/LibreLib-back/internal/config"
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

	r := routers.Api()

	// http.HandleFunc("/set-name", setNameHandler)
	// http.HandleFunc("/get-name", getNameHandler)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}
