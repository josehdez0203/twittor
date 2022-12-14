package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twittor/bd"
)

func GetTweets(w http.ResponseWriter, r *http.Request){
  ID := r.URL.Query().Get("id")
  if len(ID) <1{
    http.Error(w, "Debe mandar el parametro id", http.StatusBadRequest)
    return
  }

  if len(r.URL.Query().Get("pagina"))<1 {
    http.Error(w, "Debe mandar el parametro página", http.StatusBadRequest)
    return
  }
  pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
  if err != nil{
    http.Error(w, "El parametro página debe ser mayor a 0", http.StatusBadRequest)
    return
  }

  pag :=int64(pagina)

  respuesta, correcto := bd.GetTweets(ID, pag)
  if !correcto{
    http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
    return
  }


  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(respuesta)
}
