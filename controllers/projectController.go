package controllers

import (
	"net/http"
)

type ProjectController struct{}

func (pc ProjectController) GetAll(w http.ResponseWriter, r *http.Request) {}

func (pc ProjectController) GetId(w http.ResponseWriter, r *http.Request) {}
