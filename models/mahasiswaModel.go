package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Nama     string             `bson:"nama" json:"nama"`
	NRP      string             `bson:"nrp" json:"nrp"`
	Semester string             `bson:"semester" json:"semester"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}
