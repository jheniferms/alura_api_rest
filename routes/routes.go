package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jheniferms/alura_api_rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAll).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.Get).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
