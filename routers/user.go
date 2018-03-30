package routers

import (
	"github.com/gorilla/mux"
	"TravelShipper/controllers"
	"TravelShipper/common"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	router.HandleFunc("/users/activate", controllers.Activate).Methods("POST")
	router.HandleFunc("/users/resendcode", controllers.ResendActivateCode).Methods("POST")
	router.HandleFunc("/users/requestresetpassword", controllers.RequestResetPassword).Methods("POST")
	router.HandleFunc("/users/resetpassword", controllers.ResetPassword).Methods("POST")

	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/users/myprofile", controllers.GetMyProfile).Methods("GET")
	userRouter.HandleFunc("/users/updateprofile", controllers.UpdateProfile).Methods("POST")
	userRouter.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.PathPrefix("/users").Handler(common.AuthorizeRequest(userRouter))
	return router
}
