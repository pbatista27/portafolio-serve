package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	//init router
	router := mux.NewRouter().StrictSlash(false)

	//router inicio
	router = SetInicioRouter(router)

	//router user
	router = SetUserRouter(router)

	//router project
	router = SetProjectRouter(router)

	//router Redes sosciales
	router = SetRedesSocialesRouter(router)

	//router lenguaje
	router = SetLenguajeRouter(router)

	return router
}
