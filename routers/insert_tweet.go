package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twittor/bd"
	"twittor/models"
)

func InsertarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos invalidos"+err.Error(), 400)
		return
	}
	registro := models.CrearTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertarTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al insertar el tweet", 400)
		return

	}
	if !status {
		http.Error(w, "No se pudo insertar el tweet", 400)
		return

	}
	w.WriteHeader(http.StatusCreated)
}
