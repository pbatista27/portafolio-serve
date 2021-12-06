package controllers

import (
	"net/http"
)

type InicioController struct{}

func (ic InicioController) Inicio(w http.ResponseWriter, r *http.Request) {}

func (ic InicioController) GetContacto(w http.ResponseWriter, r *http.Request) {}

func (ic InicioController) RegisterContacto(w http.ResponseWriter, r *http.Request) {}

func (ic InicioController) ReadContacto(w http.ResponseWriter, r *http.Request) {}

func (ic InicioController) ResponseContacto(w http.ResponseWriter, r *http.Request) {}

func (ic InicioController) Curriculum(w http.ResponseWriter, r *http.Request) {}

func (ic InicioController) EditCurriculum(w http.ResponseWriter, r *http.Request) {}
