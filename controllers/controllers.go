package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jheniferms/alura_api_rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
