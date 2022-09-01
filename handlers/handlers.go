package handlers

import (
	"log"
	"net/http"
	"os"
	"twittor/middleware"
	"twittor/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {

	router := mux.NewRouter()
	router.HandleFunc("/registro", middleware.CheckBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckBD(routers.Profile)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
