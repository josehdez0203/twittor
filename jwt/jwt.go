package jwt

import (
	"fmt"
	"time"
	"twittor/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerarJWT(t models.Usuario) (string, error) {
	miClave := []byte("josehernandezc@gmail.com")

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"avatar":          t.Avatar,
		"sitioweb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		fmt.Println("🆎 " +tokenStr)
		return tokenStr, err
	}
	return tokenStr, nil
}
