package routes

import (
	"log"
	"net/http"

	"github.com/jheniferms/alura_api_rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.GetAll)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
