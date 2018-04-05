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
	locationRouter.HandleFunc("/items/{id}", controllers.ItemDetail).Methods("GET")
	router.PathPrefix("/items").Handler(common.AuthorizeRequest(locationRouter))
	return router
}