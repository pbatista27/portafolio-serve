package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pbatista27/portafolio-serve/common"
	"github.com/pbatista27/portafolio-serve/common/helpers"
	"github.com/pbatista27/portafolio-serve/data"
	"github.com/pbatista27/portafolio-serve/models"
)

type UserController struct{}

func (uc UserController) Register(w http.ResponseWriter, r *http.Request) {

	var u models.User
	var hash []byte
	var resp map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "hubo un error "+err.Error(), http.StatusBadRequest)
		return
	}

	if val, err := helpers.ValidarCorreoExiste(u.Email); err != nil {
		http.Error(w, "Ocurrio un error en el servidor", http.StatusBadRequest)
		return
	} else {

		if val {
			http.Error(w, "error este usuario ya existe", http.StatusUnauthorized)
			return
		}
	}

	token, err := common.NewToken().Generar(u)

	if err != nil {
		http.Error(w, "error al generar token "+err.Error(), http.StatusBadRequest)
		return
	}

	hash, err = common.NewBcrypt().Encriptar(u.Password)
	if err != nil {
		http.Error(w, "ha ocurrido un error en el sevidor "+err.Error(), http.StatusBadRequest)
		return
	}

	u.Password = string(hash)
	err = data.UserRepository{}.CrearUsuario(&u)

	if err != nil {
		http.Error(w, "error al registrar usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	resp = map[string]interface{}{
		"name":  u.Name,
		"email": u.Email,
		"jwt":   token,
		"ok":    true,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("type-Content", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request) {

	var login models.Login
	var mu models.User
	var resp map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}

	valid, err := helpers.ValidarCorreoExiste(login.Email)

	if err != nil {
		http.Error(w, "ocurrio un error al iniciar sesion", http.StatusBadRequest)
		return
	}

	if !valid {
		http.Error(w, "email o contraseña invalido", http.StatusUnauthorized)
		return
	}

	mu, valid, err = data.UserRepository{}.Login(login.Email, login.Password)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}

	if !valid {
		http.Error(w, "email o contraseña invalido", http.StatusUnauthorized)
		return
	}

	jwt, err := common.NewToken().Generar(mu)

	if err != nil {
		http.Error(w, "error al generar token "+err.Error(), http.StatusBadRequest)
		return
	}

	resp = map[string]interface{}{
		"name":  mu.Name,
		"email": mu.Email,
		"jwt":   jwt,
		"ok":    true,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
