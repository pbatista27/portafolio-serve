package controllers

import (
	"net/http"
)

type Inicio struct{}

func (ic Inicio) Inicio(w http.ResponseWriter, r *http.Request) {}

func (ic Inicio) GetContacto(w http.ResponseWriter, r *http.Request) {}

func (ic Inicio) RegisterContacto(w http.ResponseWriter, r *http.Request) {}

func (ic Inicio) ReadContacto(w http.ResponseWriter, r *http.Request) {}

func (ic Inicio) ResponseContacto(w http.ResponseWriter, r *http.Request) {}

func (ic Inicio) Curriculum(w http.ResponseWriter, r *http.Request) {}

func (ic Inicio) EditCurriculum(w http.ResponseWriter, r *http.Request) {}
