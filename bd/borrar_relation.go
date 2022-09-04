package bd

import (
	"context"
	"time"
	"twittor/models"
)

func BorrarRelation(t models.Relation)(bool, error){
  ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

  defer cancel()
  db:= MongoCN.Database("twittor")
  col:= db.Collection("relations")

  _, err := col.DeleteOne(ctx, t)
  if err !=nil{
    return false, err
  }
  return true, nil
}
