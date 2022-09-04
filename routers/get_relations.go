package routers

import (
	"encoding/json"
	"net/http"
	"twittor/bd"
	"twittor/models"
)


func GetRelations(w http.ResponseWriter, r *http.Request){
  ID := r.URL.Query().Get("id")
  var t models.Relation
  t.UsuarioID =IDUsuario
  t.UsuarioRelationID=ID

  var resp models.RespuestaRelation

  status, err := bd.GetRelations(t)
  if err !=nil || !status{
    resp.Status = false
  }else{
    resp.Status= true
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(resp)
}
