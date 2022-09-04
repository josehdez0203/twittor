package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RespuestaTweetsSeguidores struct{
  ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
  UsuarioID string `bson:"usuarioid" json:"userId,omitempty"`
  UsuarioRelationID string `bson:"usuariorelationid" json:"userRelationId,omitempty"`
  Tweet struct{ 
    Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
    Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
    ID string `_id:"mensaje" json:"_id,omitempty"`
    }
}
