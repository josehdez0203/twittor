package bd

import (
	"context"
	"time"
	"twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarTweet(t models.CrearTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

  result, err := col.InsertOne(ctx, registro)
  if err!=nil{
    return "", false, err
  }
  objID, _ := result.InsertedID.(primitive.ObjectID)

  return objID.String(), true, nil
}
