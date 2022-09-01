package routers

import(
  "errors"
  "strings"
  jwt "github.com/dgrijalva/jwt-go"
  "twittor/bd"
  "twittor/models"
)

var Email string
var IDUsuario string

func ProcesarToken(tk string)(*models.Claim, bool, string, error){
  miClave :=[]byte("josehernandezc@gmail.com")

  claims := &models.Claim{}

  splitToken := strings.Split(tk, "Bearer")
  if len(splitToken) !=2{
    return claims, false, string(""), errors.New("formato de token invalido")
  }
  tk= strings.TrimSpace(splitToken[1])

  tkNew, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
    return miClave, nil
  })
  if err ==nil{
    _, encontrado, _ :=bd.CheckExistingUser(claims.Email)
    if encontrado{
    Email=claims.Email
    IDUsuario= claims.ID.Hex()
    }
    return claims, encontrado, IDUsuario, nil
  }
  if !tkNew.Valid{
    return claims, false, string(""), errors.New("Token invalido")
  }

  return claims, false, string(""), err
}

