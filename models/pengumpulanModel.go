package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BerkasPengumpulan struct {
	NamaBerkas string `bson:"nama_berkas" json:"nama_berkas"`
	URLBerkas  string `bson:"url_berkas" json:"url_berkas"`
}

type Pengumpulan struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	Nama          string             `bson:"nama" json:"nama"`
	NomorPengenal string             `bson:"nomor_pengenal" json:"nomor_pengenal"`
	Jenis         string             `bson:"jenis" json:"jenis"`
	Berkas        BerkasPengumpulan  `bson:"berkas" json:"berkas"`
	Tgl           string             `bson:"tgl" json:"tgl"`
}
