package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetInicioRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/", controllers.Inicio{}.Inicio).Methods("GET")
	router.HandleFunc("/contacto", controllers.Inicio{}.RegisterContacto).Methods("POST")
	router.HandleFunc("/contacto", controllers.Inicio{}.GetContacto).Methods("GET")
	router.HandleFunc("/contacto/{id}", controllers.Inicio{}.ReadContacto).Methods("GET")
	router.HandleFunc("/contacto/response", controllers.Inicio{}.ResponseContacto).Methods("POST")

	return router
}
