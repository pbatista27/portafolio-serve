package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetInicioRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/", controllers.InicioController{}.Inicio).Methods("GET")
	router.HandleFunc("/contacto", controllers.InicioController{}.RegisterContacto).Methods("POST")
	router.HandleFunc("/contacto", controllers.InicioController{}.GetContacto).Methods("GET")
	router.HandleFunc("/contacto/{id}", controllers.InicioController{}.ReadContacto).Methods("GET")
	router.HandleFunc("/contacto/response", controllers.InicioController{}.ResponseContacto).Methods("POST")

	return router
}
