package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Nama     string             `bson:"nama" json:"nama"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}
