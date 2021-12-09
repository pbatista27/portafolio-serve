package controllers

import (
	"net/http"
)

type Project struct{}

func (pc Project) GetAll(w http.ResponseWriter, r *http.Request) {}

func (pc Project) GetId(w http.ResponseWriter, r *http.Request) {}
