package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty" validate:"required"`
	NRP      string             `bson:"nrp,omitempty" json:"nrp,omitempty" validate:"required"`
	Semester string             `bson:"semester,omitempty" json:"semester,omitempty" validate:"required"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty" validate:"required"`
	Password string             `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
}
