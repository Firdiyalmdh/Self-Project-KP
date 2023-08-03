package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BerkasPermohonan struct {
	NamaBerkas string `bson:"nama_berkas" json:"nama_berkas"`
	URLBerkas  string `bson:"url_berkas" json:"url_berkas"`
}

type Pemohon struct {
	Nama           string `bson:"nama" json:"nama"`
	Nomor_Pengenal string `bson:"nomor_pengenal" json:"nomor_pengenal"`
}

type Permohonan struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Tipe     string             `bson:"tipe" json:"tipe"`
	Status   string             `bson:"status" json:"status"`
	Pemohon  Pemohon            `bson:"pemohon" json:"pemohon"`
	Berkas   BerkasPermohonan   `bson:"berkas" json:"berkas"`
	TglMasuk string             `bson:"tgl_masuk" json:"tgl_masuk"`
	Tujuan   string             `bson:"tujuan" json:"tujuan"`
	Hasil    string             `bson:"hasil" json:"hasil"`
}
