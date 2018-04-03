package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	//router = SetTestingRouter(router)

	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Bookmark entity
	router = SetBookmarkRoutes(router)
	router = SetLocationRoutes(router)

	return router
}