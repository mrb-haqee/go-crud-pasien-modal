package sql

import (
	"fmt"

	"github.com/mrb-haqee/go-crud-pasien-model/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connectDB struct {
	db *gorm.DB
}

func NewDB() *connectDB {
	return &connectDB{}
}

func (p *connectDB) Connect(dbname model.DbName) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbname.Host, dbname.Username, dbname.Password, dbname.DatabaseName, dbname.Port)
	dbconn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return dbconn, err
}
