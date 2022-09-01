package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"twittor/bd"
	"twittor/jwt"
	"twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña no válidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.Login(t.Email, t.Password)
		fmt.Println(documento)
	if !existe {
		http.Error(w, "Usuario y/o contraseña inválidos", 400)
		return
	}
	jwtKey, err := jwt.GenerarJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al generar el token "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: jwtKey, Expires: expirationTime})
}
