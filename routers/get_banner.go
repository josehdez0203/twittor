package routers

import (
	"io"
	"log"
	"net/http"
	"os"
	"twittor/bd"
)


func GetBanner(w http.ResponseWriter, r *http.Request){
  ID := r.URL.Query().Get("id")

  if len(ID) <1{
    http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
    return
  }
  perfil, err := bd.UsuarioXID(ID)
  if err!=nil{
 http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
    return
  }
log.Println("🆚 " +perfil.Banner)

  OpenFile,miErr  := os.Open("uploads/banners/"+perfil.Banner)
  if miErr!=nil{
     http.Error(w, "imagen no encontrada", http.StatusBadRequest)
    return
  }
  _, er := io.Copy(w, OpenFile)
  if er !=nil{
http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
    return
  }
}
