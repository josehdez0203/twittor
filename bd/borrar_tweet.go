package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func BorrarTweet(ID string, UserID string) error{
  ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")
  objID,_ :=primitive.ObjectIDFromHex(ID)
  condicion:=bson.M{
    "_id": objID,
    "userid":UserID,
  }
  _, err := col.DeleteOne(ctx, condicion)
  return err
}
