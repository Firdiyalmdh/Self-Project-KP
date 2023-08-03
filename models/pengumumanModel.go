package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Announcer struct {
	Nama           string `bson:"nama" json:"nama" validate:"required"`
	Nomor_Pengenal string `bson:"nomor_pengenal" json:"nomor_pengenal" validate:"required"`
}

type Content struct {
	Tgl  string `bson:"tgl" json:"tgl" validate:"required"`
	Data string `bson:"data" json:"data" validate:"required"`
}

type Pengumuman struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Jenis     string             `bson:"jenis" json:"jenis" validate:"required"`
	Announcer Announcer          `bson:"announcer" json:"announcer" validate:"required"`
	Content   Content            `bson:"content" json:"content" validate:"required"`
}
