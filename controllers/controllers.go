package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jheniferms/alura_api_rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	personalidades := models.GetAll()

	json.NewEncoder(w).Encode(personalidades)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := getIdFromUrl(r)

	personalidade := models.Get(id)

	json.NewEncoder(w).Encode(personalidade)
}

func getIdFromUrl(r *http.Request) string {
	vars := mux.Vars(r)

	return vars["id"]
}

func Create(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	novaPersonalidade = models.Create(novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := getIdFromUrl(r)

	personalidade := models.Delete(id)

	json.NewEncoder(w).Encode(personalidade)
}

func Update(w http.ResponseWriter, r *http.Request) {

	id := getIdFromUrl(r)

	personalidadeUpdate := models.Get(id)

	json.NewDecoder(r.Body).Decode(&personalidadeUpdate)

	models.Update(personalidadeUpdate)

	json.NewEncoder(w).Encode(personalidadeUpdate)
}
