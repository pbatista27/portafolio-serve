package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetProjectRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/project/", controllers.ProjectController{}.GetAll).Methods("GET")
	router.HandleFunc("/project/{id}", controllers.ProjectController{}.GetId).Methods("GET")
	return router

}
