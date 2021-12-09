package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/data"
	"github.com/pbatista27/portafolio-serve/models"
)

type Redes struct{}

func (rs Redes) GetAll(w http.ResponseWriter, r *http.Request) {

	redes, err := data.RedesRepository{}.GetAll()

	if err != nil {
		http.Error(w, "ocurrio un error al obtener la redes sociales "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(redes)

}
func (rs Redes) Show(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	redes, err := data.RedesRepository{}.Show(id)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(redes)
}

func (rs Redes) Create(w http.ResponseWriter, r *http.Request) {

	var redes models.SocialNetworks

	err := json.NewDecoder(r.Body).Decode(&redes)

	if err != nil {
		http.Error(w, "error al recibir datos "+err.Error(), http.StatusBadRequest)
		return
	}

	err = data.RedesRepository{}.Create(redes)

	if err != nil {
		http.Error(w, "ocurri un error "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(redes)

}

func (rs Redes) Update(w http.ResponseWriter, r *http.Request) {

	var redes models.SocialNetworks

	if err := json.NewDecoder(r.Body).Decode(&redes); err != nil {
		http.Error(w, "ocurrio un error al obtener los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	err := data.RedesRepository{}.Update(redes)

	if err != nil {
		http.Error(w, "ocurrio un error al actualizar "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (rs Redes) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	err := data.RedesRepository{}.Delete(id)

	if err != nil {
		http.Error(w, "ocurri un error al eliminar "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Contet-type", "Application/json")

}
