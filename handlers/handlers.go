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
	router.HandleFunc("/profile", middleware.CheckBD(middleware.ValidarJWT(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/modProfile", middleware.CheckBD(middleware.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/insertarTweet", middleware.CheckBD(middleware.ValidarJWT(routers.InsertarTweet))).Methods("POST")
	router.HandleFunc("/getTweets", middleware.CheckBD(middleware.ValidarJWT(routers.GetTweets))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
