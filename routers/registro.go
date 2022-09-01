package routers

import (
	"encoding/json"
	"net/http"
	"twittor/bd"
	"twittor/models"
)

//Registro de usuarios
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El tamaño del Password es menos de 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.UsuarioExistente(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese Email", 400)
		return
	}

	_, status, err := bd.InsertarUsuario(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar registrar usuario"+err.Error(), 400)
		return
	}
	if status {
		http.Error(w, "No se pudo registrar usuario"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
