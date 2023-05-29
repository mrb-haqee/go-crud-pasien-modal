package model

import "gorm.io/gorm"

type DbName struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Pasien struct {
	gorm.Model
	NamaLengkap  string `gorm:"type:varchar" json:"nama_lengkap"`
	NIK          string `gorm:"type:varchar" json:"nik"`
	Gender       string `gorm:"type:varchar(1)" json:"gender"`
	Umur         uint64 `gorm:"type:smallint" json:"umur"`
	TempatLahir  string `gorm:"type:varchar" json:"tempat_lahir"`
	TanggalLahir string `gorm:"type:varchar" json:"tanggal_lahir"`
	NoHP         int    `gorm:"type:bigint" json:"no_hp"`
	Alamat       string `gorm:"type:varchar" json:"alamat"`
	Keluhan      string `gorm:"type:text" json:"keluhan"`
}
