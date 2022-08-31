package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimento" json:"fechaNacimento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password"`
	Avatar          string             `bson:"avatar" json:"avatar"`
	Banner          string             `bson:"banner" json:"banner"`
	Biografia       string             `bson:"biografia" json:"biografia"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb"`
}
