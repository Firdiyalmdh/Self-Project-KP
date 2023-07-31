package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty" validate:"required"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty" validate:"required"`
	Password string             `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
}
