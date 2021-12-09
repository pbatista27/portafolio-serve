package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetLenguajeRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/lenguaje", controllers.Lenguaje{}.GetAll).Methods("GET")
	router.HandleFunc("/lenguaje/{id}", controllers.Lenguaje{}.Show).Methods("GET")
	router.HandleFunc("/lenguaje", controllers.Lenguaje{}.Create).Methods("POST")
	router.HandleFunc("/lenguaje", controllers.Lenguaje{}.Update).Methods("PUT")
	router.HandleFunc("/lenguaje/{id}", controllers.Lenguaje{}.Delete).Methods("DELETE")

	return router

}
