package bd

import (
	"context"
	"fmt"
	"log"
	"time"
	"twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscarPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": objID}
  err := col.FindOne(ctx, condicion).Decode(&perfil)
  perfil.Password=""

  if err!=nil{
    log.Println("Respuesta fallida 📛" )
    fmt.Println("Registro no encontrado "+err.Error())
    return perfil, err
  }
  log.Println("Respuesta Ok, 🆗" )
  return perfil, nil

}
