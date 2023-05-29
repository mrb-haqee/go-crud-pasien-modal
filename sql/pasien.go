package sql

import (
	"github.com/mrb-haqee/go-crud-pasien-model/model"
	"gorm.io/gorm"
)

type PasienDB interface {
	FindAll(pasiens *[]model.Pasien) error
	FindId(id int, pasien *model.Pasien) error
	Add(pasien model.Pasien) error
	Update(id int, pasien model.Pasien) error
	Delete(id int) error
}

type pasienDB struct {
	db *gorm.DB
}

func NewPasien(db *gorm.DB) *pasienDB {
	return &pasienDB{db: db}
}

func (p *pasienDB) FindAll(pasiens *[]model.Pasien) error {
	return p.db.Find(pasiens).Error
}

func (p *pasienDB) FindId(id int, pasien *model.Pasien) error {
	return p.db.First(pasien, "id = ?", id).Error
}

func (p *pasienDB) Add(pasien model.Pasien) error {
	return p.db.Create(&pasien).Error
}

func (p *pasienDB) Update(id int, pasien model.Pasien) error {
	return p.db.Model(model.Pasien{}).Where("id = ?", id).Updates(pasien).Error
}

func (p *pasienDB) Delete(id int) error {
	return p.db.Where("id = ?", id).Delete(&model.Pasien{}).Error
}
