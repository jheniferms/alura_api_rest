package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jheniferms/alura_api_rest/database"
	"github.com/jheniferms/alura_api_rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade

	database.DB.Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := getIdFromUrl(r)

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func getIdFromUrl(r *http.Request) string {
	vars := mux.Vars(r)

	return vars["id"]
}

func Create(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := getIdFromUrl(r)

	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func Update(w http.ResponseWriter, r *http.Request) {

	id := getIdFromUrl(r)

	var personalidadeUpdate models.Personalidade

	database.DB.Find(&personalidadeUpdate, id)

	json.NewDecoder(r.Body).Decode(&personalidadeUpdate)

	database.DB.Save(&personalidadeUpdate)

	json.NewEncoder(w).Encode(personalidadeUpdate)
}
