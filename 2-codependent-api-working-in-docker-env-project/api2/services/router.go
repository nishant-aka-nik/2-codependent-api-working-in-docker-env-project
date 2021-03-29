package services

import (
	"log"
	"net/http"
	"sellerapi2/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InitServer provides the routes for the API and then starts the server
func InitServer() {
	router := mux.NewRouter()

	// routes are routing to the controllers
	router.HandleFunc("/getapidata", controllers.GetAPIData).Methods("POST")

	// Starting server and handling CORS
	port := ":8081"
	log.Println("Starting router at port " + port + " with uri : /getapidata")
	http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router))
}
