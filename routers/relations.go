package routers

import (
	"net/http"
	"twittor/bd"
	"twittor/models"
)

func InsertRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "EL parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelationID = ID

	status, err := bd.InsertarRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al insertar relación", http.StatusBadRequest)
		return
	}
  if !status{
	http.Error(w, "No se pudo insertar la relación", http.StatusBadRequest)
		return
  }
  w.WriteHeader(http.StatusCreated)
}
