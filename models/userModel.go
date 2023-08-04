package models

type User struct {
	Id            string `bson:"_id" json:"_id"`
	Email         string `bson:"email" json:"email"`
	Nama          string `bson:"nama" json:"nama"`
	NomorPengenal string `bson:"nomor_pengenal" json:"nomor_pengenal"`
	Role          string `bson:"role" json:"role"`
	Login         string `bson:"login" json:"login"`
}
