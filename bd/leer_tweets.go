package bd

import (
	"context"
	"time"
	"twittor/models"

	"go.mongodb.org/mongo-driver/bson"
)

func LeerTweets(ID string, pagina int) ([]models.RespuestaTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relations")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "usuariorelationid",
			"foreignField": "userid",
			"as":           "tweets",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweets"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweets.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condiciones)

	var result []models.RespuestaTweetsSeguidores

	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
