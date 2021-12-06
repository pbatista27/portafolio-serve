package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetUserRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/user/registro", controllers.UserController{}.Register).Methods("POST")
	router.HandleFunc("/user/login", controllers.UserController{}.Login).Methods("POST")

	return router
}
