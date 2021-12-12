package routers

import (
	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/controllers"
)

func SetStudyRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/estudio", controllers.Study{}.GetAll).Methods("GET")
	router.HandleFunc("/estudio/{id}", controllers.Study{}.Show).Methods("GET")
	router.HandleFunc("/estudio", controllers.Study{}.Create).Methods("POST")
	router.HandleFunc("/estudio", controllers.Study{}.Update).Methods("PUT")
	router.HandleFunc("/estudio/{id}", controllers.Study{}.Delete).Methods("DELETE")

	return router
}
