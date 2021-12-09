package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/data"
	"github.com/pbatista27/portafolio-serve/models"
)

type Lenguaje struct{}

func (l Lenguaje) GetAll(w http.ResponseWriter, r *http.Request) {

	langs, err := data.LanguajeRespository{}.GetAll()

	if err != nil {
		http.Error(w, "Error al obtener los lenguajes "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(langs)

}

func (l Lenguaje) Create(w http.ResponseWriter, r *http.Request) {

	var lang models.Languaje

	err := json.NewDecoder(r.Body).Decode(&lang)

	if err != nil {
		http.Error(w, "error al recibir data "+err.Error(), http.StatusBadRequest)
		return
	}

	data, err := data.LanguajeRespository{}.Create(lang)

	if err != nil {
		http.Error(w, "error al registrar lenguaje "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "Application/json")
	json.NewEncoder(w).Encode(data)

}

func (l Lenguaje) Show(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	lang, err := data.LanguajeRespository{}.Show(id)

	if err != nil {
		http.Error(w, "error al mostrar lenguaje "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "Application/json")
	json.NewEncoder(w).Encode(lang)

}

func (l Lenguaje) Update(w http.ResponseWriter, r *http.Request) {

	var lang models.Languaje

	err := json.NewDecoder(r.Body).Decode(&lang)

	if err != nil {
		http.Error(w, "ocurrio un error al recibir los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	data, err := data.LanguajeRespository{}.Update(lang)

	if err != nil {
		http.Error(w, "ocurrio un erro al actualizar datos "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "Application/json")
	json.NewEncoder(w).Encode(data)

}

func (l Lenguaje) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	valid, err := data.LanguajeRespository{}.Delete(id)

	if err != nil {
		http.Error(w, "Error al eliminar lenguaje "+err.Error(), http.StatusBadRequest)
		return
	}

	if !valid {
		http.Error(w, "Error al eliminar lenguaje", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "Application/json")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}
