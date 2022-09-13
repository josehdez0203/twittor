package middleware

import (
	"log"
	"net/http"
	"twittor/routers"
)


func ValidarJWT(next http.HandlerFunc) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request){
    _,_,_, err := routers.ProcesarToken(r.Header.Get("Authorization"))
    if err != nil{
      log.Println( "🧲 ⚠️ :"+ "token invalido")
      http.Error(w, "Error en token! "+err.Error(), http.StatusBadRequest)
      return
    }
    next.ServeHTTP(w,r)
  }
}
