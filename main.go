package main

import (
	"log"
	"twittor/bd"
	"twittor/handlers"
)
func main(){
  if bd.CheckConnection()==0{
    log.Fatal("Sin conexión a la BD")
    return
  }
  handlers.Manejadores()
}
