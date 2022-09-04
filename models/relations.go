package models

type Relation struct{
  UsuarioID string `bson:"usuarioid" json:"usuarioId"`
  UsuarioRelationID string `bson:"usuariorelationid" json:"usuarioRelationId"`
}
 type RespuestaRelation struct{
  Status bool `json:"status"`
 }
