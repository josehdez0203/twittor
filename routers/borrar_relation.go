package routers

import (
	"net/http"
	"twittor/bd"
	"twittor/models"
)

func BorraRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelationID = ID

	status, err := bd.BorrarRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la relación", http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se logró borrar la relación", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
