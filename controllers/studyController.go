package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pbatista27/portafolio-serve/data"
	"github.com/pbatista27/portafolio-serve/models"
)

type Study struct{}

func (s Study) GetAll(w http.ResponseWriter, r *http.Request) {

	study, err := data.StudyRepository{}.GelAll()

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}
	resp := map[string]interface{}{
		"ok":     true,
		"studio": study,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (s Study) Show(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	study, err := data.StudyRepository{}.Show(id)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}
	resp := map[string]interface{}{
		"ok":      true,
		"estudio": study,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (s Study) Create(w http.ResponseWriter, r *http.Request) {

	var study models.Study

	err := json.NewDecoder(r.Body).Decode(&study)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}

	data, err := data.StudyRepository{}.Create(study)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"ok":      true,
		"estudio": data,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (s Study) Update(w http.ResponseWriter, r *http.Request) {

	var study models.Study

	err := json.NewDecoder(r.Body).Decode(&study)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}

	err = data.StudyRepository{}.Update(study)
	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}

func (s Study) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	err := data.StudyRepository{}.Delete(id)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}
