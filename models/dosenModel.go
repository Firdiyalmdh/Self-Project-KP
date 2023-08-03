package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dosen struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Nama     string             `bson:"nama" json:"nama"`
	NIP      string             `bson:"nip" json:"nip"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}
