package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	IdUser        string `bson:"id" json:"json"`
	Email         string `bson:"email" json:"email"`
	Nama          string `bson:"nama" json:"nama"`
	NomorPengenal string `bson:"nomor_pengenal" json:"nomor_pengenal"`
	Role          string `bson:"role" json:"role"`
	Login         string `bson:"login" json:"login"`
}
