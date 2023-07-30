package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Berkas struct {
	NamaBerkas string `bson:"nama_berkas,omitempty" json:"nama_berkas,omitempty"`
	URLBerkas  string `bson:"url_berkas,omitempty" json:"url_berkas,omitempty"`
}

type Pemohon struct {
	Nama           string `bson:"nama,omitempty" json:"nama,omitempty"`
	Nomor_Pengenal string `bson:"nomor_pengenal,omitempty" json:"nomor_pengenal,omitempty"`
}

type Permohonan struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis    string             `bson:"jenis,omitempty" json:"jenis,omitempty"`
	Status   string             `bson:"status,omitempty" json:"status,omitempty"`
	Pemohon  Pemohon            `bson:"pemohon,omitempty" json:"pemohon,omitempty"`
	Tujuan   string             `bson:"tujuan,omitempty" json:"tujuan,omitempty"`
	Berkas   Berkas             `bson:"berkas,omitempty" json:"berkas,omitempty"`
	TglMasuk string             `bson:"tgl_masuk,omitempty" json:"tgl_masuk,omitempty"`
}
