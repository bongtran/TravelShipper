package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = SetTestingRouter(router)

	router = SetStellarRouters(router)

	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Bookmark entity
	router = SetBookmarkRoutes(router)

	return router
}