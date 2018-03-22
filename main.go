package main

import (
	"TravelShipper/common"
	"net/http"
	"TravelShipper/routers"
	"log"
)

func main() {
	common.StartUp()

	router := routers.InitRouters()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("listerning...")

	server.ListenAndServe()
}
