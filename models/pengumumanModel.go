package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Announcer struct {
	Nama           string `bson:"nama,omitempty" json:"nama,omitempty" validate:"required"`
	Nomor_Pengenal string `bson:"nomor_pengenal,omitempty" json:"nomor_pengenal,omitempty" validate:"required"`
}

type Content struct {
	Tgl  string `bson:"tgl,omitempty" json:"tgl,omitempty" validate:"required"`
	Data string `bson:"data,omitempty" json:"data,omitempty" validate:"required"`
}

type Pengumuman struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis     string             `bson:"jenis,omitempty" json:"jenis,omitempty" validate:"required"`
	Announcer Announcer          `bson:"announcer,omitempty" json:"announcer,omitempty" validate:"required"`
	Content   Content            `bson:"content,omitempty" json:"content,omitempty" validate:"required"`
}
