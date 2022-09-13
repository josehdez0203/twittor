package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"twittor/bd"
)

func GetProfile(w http.ResponseWriter, r *http.Request){
  ID := r.URL.Query().Get("id")
  if len(ID) <1 {
    http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
    return
  }

  perfil, err := bd.BuscarPerfil(ID)
  log.Println("Request Ok, inicia respuesta 🆗" )
  if err!=nil{

    http.Error(w,"Ocurrió un error al intentar buscar registro " +err.Error(), 400)
  }

  w.Header().Set("Content-Type", "applicacion/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(perfil)
}
