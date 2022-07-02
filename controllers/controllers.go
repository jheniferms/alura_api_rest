package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jheniferms/alura_api_rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
