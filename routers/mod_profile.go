package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"twittor/bd"
	"twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
    log.Println("🎅 Error de datos ", )
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificarUsuario(t, IDUsuario)
	if err != nil {
  log.Println("🎅 Sucedío error al modificar BD" )
		http.Error(w, "Ocurrió error al modificar el usuario "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se pudo modificar el usuario ", 400)
	}
	log.Println("🆚 Se actualizó usuario", IDUsuario)
	w.WriteHeader(http.StatusCreated)
}
