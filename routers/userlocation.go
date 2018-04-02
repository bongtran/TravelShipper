package routers

import (
	"github.com/gorilla/mux"
	"TravelShipper/controllers"
	"TravelShipper/common"
)

func SetLocationRoutes(router *mux.Router) *mux.Router {
	locationRouter := mux.NewRouter()
	locationRouter.HandleFunc("/locations/mylocation", controllers.GetMyLocation).Methods("GET")
	locationRouter.HandleFunc("/locations/mylocation", controllers.SetLocation).Methods("POST")
	//locationRouter.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.PathPrefix("/locations").Handler(common.AuthorizeRequest(locationRouter))
	return router
}