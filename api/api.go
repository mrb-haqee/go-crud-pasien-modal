package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrb-haqee/go-crud-pasien-model/sql"
)

type API struct {
	db  sql.PasienDB
	mux *mux.Router
}

func NewAPI(db sql.PasienDB) API {

	r := mux.NewRouter()
	api := API{db: db, mux: r}

	r.HandleFunc("/", api.Index).Methods("GET")
	r.HandleFunc("/pasien", api.Index).Methods("GET")
	r.HandleFunc("/pasien/form", api.Form).Methods("GET", "POST")
	r.HandleFunc("/pasien/store", api.Store).Methods("POST")
	r.HandleFunc("/pasien/delete", api.Delete).Methods("POST")

	return api
}

func (api *API) Handler() http.Handler {
	return api.mux
}

func (api *API) Start() error {
	log.Print("Run server at http://localhost:8080")
	return http.ListenAndServe(":8080", api.Handler())
}
