package main

import (
	"log"

	"github.com/mrb-haqee/go-crud-pasien-model/api"
	"github.com/mrb-haqee/go-crud-pasien-model/model"
	"github.com/mrb-haqee/go-crud-pasien-model/sql"
)

var DB model.DbName = model.DbName{
	Host:         "localhost",
	Username:     "postgres",
	Password:     "mrb28",
	DatabaseName: "crud_pasien",
	Port:         5432,
	Schema:       "public",
}

func main() {

	db := sql.NewDB()
	conn, err := db.Connect(DB)
	if err != nil {
		panic(err)
	}
	log.Print("Terhubung ke DB")

	conn.AutoMigrate(&model.Pasien{})

	pasienDB := sql.NewPasien(conn)

	api := api.NewAPI(pasienDB)
	api.Start()
}
