package bd

import(
  "twittor/models"

  "golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.Usuario, bool){
  usu, encontrado, _ :=UsuarioExistente(email)
  if !encontrado{ //No se encontro usuario
    return usu, false
  }
  passwordBytes := []byte(password)
  passwordBD := []byte(usu.Password)

  err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
  if err!=nil{
    return usu, false
  }
  return usu, true
}
