package routers

import (
	"github.com/gorilla/mux"
	"TravelShipper/controllers"
	"TravelShipper/common"
)

func SetLocationRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/locations/myprofile", controllers.GetMyProfile).Methods("GET")
	userRouter.HandleFunc("/locations/updateprofile", controllers.UpdateProfile).Methods("POST")
	userRouter.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.PathPrefix("/users").Handler(common.AuthorizeRequest(userRouter))
	return router
}