package routers

import (
	"github.com/gorilla/mux"
	"TravelShipper/common"
	"TravelShipper/controllers"
)

func SetItemRoutes(router *mux.Router) *mux.Router {
	locationRouter := mux.NewRouter()
	locationRouter.HandleFunc("/items", controllers.MyItems).Methods("GET")
	locationRouter.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	//locationRouter.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.PathPrefix("/locations").Handler(common.AuthorizeRequest(locationRouter))
	return router
}