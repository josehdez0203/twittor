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
	router.HandleFunc("/getTweets", middleware.CheckBD(middleware.ValidarJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middleware.CheckBD(middleware.ValidarJWT(routers.BorrarTweet))).Methods("DELETE")
	
  router.HandleFunc("/setAvatar", middleware.CheckBD(middleware.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/setBanner", middleware.CheckBD(middleware.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middleware.CheckBD(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/getAvatar", middleware.CheckBD(routers.GetAvatar)).Methods("GET")

  router.HandleFunc("/seguir", middleware.CheckBD(middleware.ValidarJWT(routers.InsertRelation))).Methods("POST")
  router.HandleFunc("/noSeguir", middleware.CheckBD(middleware.ValidarJWT(routers.BorraRelation))).Methods("DELETE")
  router.HandleFunc("/consultaRelacion", middleware.CheckBD(middleware.ValidarJWT(routers.GetRelations))).Methods("GET")

  router.HandleFunc("/getUsers", middleware.CheckBD(middleware.ValidarJWT(routers.GetAllUsers))).Methods("GET")
  router.HandleFunc("/getTweetsSeguidores", middleware.CheckBD(middleware.ValidarJWT(routers.LeerTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
