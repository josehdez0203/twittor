package routers

import (
	"io"
	"log"
	"net/http"
	"os"
	"twittor/bd"
)


func GetAvatar(w http.ResponseWriter, r *http.Request){
  ID := r.URL.Query().Get("id")

  if len(ID) <1{
    http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
    return
  }
  perfil, err := bd.UsuarioXID(ID)
  if err !=nil{
 http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
    return
  }

  OpenFile,miErr  := os.Open("uploads/avatars/"+perfil.Avatar)
  if miErr!=nil{
     http.Error(w, "imagen no encontrado", http.StatusBadRequest)
    return
  } 
  log.Println("🆚 " +perfil.Avatar)
  _, er := io.Copy(w, OpenFile)
  if er !=nil{
http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
    return
  }
}
