package bd

import (
	"context"
	"fmt"
	"time"
	"twittor/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRelations(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relations")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelationid": t.UsuarioRelationID,
	}

	var resultado models.Relation
	fmt.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
