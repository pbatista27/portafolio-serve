package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetRedesSocialesRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/redes", controllers.Redes{}.GetAll).Methods("GET")
	router.HandleFunc("/redes/{id}", controllers.Redes{}.Show).Methods("GET")
	router.HandleFunc("/redes", controllers.Redes{}.Create).Methods("POST")
	router.HandleFunc("/redes", controllers.Redes{}.Update).Methods("PUT")
	router.HandleFunc("/redes/{id}", controllers.Redes{}.Delete).Methods("DELETE")

	return router
}
