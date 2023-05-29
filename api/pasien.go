package api

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/mrb-haqee/go-crud-pasien-model/model"
)

func (api *API) Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"data": template.HTML(api.GetData()),
	}

	temp, _ := template.ParseFiles("view/home.html")
	temp.Execute(w, data)

}

func (api *API) GetData() string {

	buffer := bytes.Buffer{}

	var pasiens []model.Pasien
	err := api.db.FindAll(&pasiens)
	if err != nil {
		panic(err)
	}

	inc := template.FuncMap{
		"inc": func(a, b int) int {
			return a + b
		},
	}

	data := map[string]any{
		"data": pasiens,
	}

	temp, _ := template.New("pasien.html").Funcs(inc).ParseFiles("view/pasien.html")

	temp.ExecuteTemplate(&buffer, "pasien.html", data)

	return buffer.String()
}

// drop modal
func (api *API) Form(w http.ResponseWriter, r *http.Request) {
	IdStr := r.URL.Query().Get("id")
	if IdStr != "" {
		id, _ := strconv.Atoi(IdStr)
		var pasien model.Pasien

		api.db.FindId(id, &pasien)

		data := map[string]any{
			"data": pasien,
		}

		temp, err := template.ParseFiles("view/form.html")
		if err != nil {
			http.Error(w, "ini error", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return
	}

	temp, err := template.ParseFiles("view/form.html")
	if err != nil {
		http.Error(w, "ini error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func (api *API) Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		log.Print("masuk store")
		r.ParseForm()

		var pasien model.Pasien

		pasien.NamaLengkap = r.FormValue("name_lengkap")
		pasien.NIK = r.FormValue("nik")
		pasien.Gender = r.FormValue("gender")
		pasien.Umur, _ = strconv.ParseUint(r.FormValue("umur"), 10, 64)
		pasien.TempatLahir = r.FormValue("tempat_lahir")
		pasien.TanggalLahir = r.FormValue("tanggal_lahir")
		pasien.NoHP, _ = strconv.Atoi(r.FormValue("no_hp"))
		pasien.Alamat = r.FormValue("alamat")
		pasien.Keluhan = r.FormValue("keluhan")

		if api.IsStructEmpty(pasien) {
			data := map[string]any{
				"message": "Data Gaboleh ada yang kosong yaa!",
			}
			RespJson(w, http.StatusOK, data)
			return

		}

		log.Printf("%+v", pasien)

		idStr := r.FormValue("id")
		data := make(map[string]any)

		if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			log.Println(id)
			err := api.db.Update(id, pasien)
			if err != nil {
				RespError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]any{
				"message": "Data berhasil Diupdate",
				"data":    template.HTML(api.GetData()),
			}

		} else {

			err := api.db.Add(pasien)
			if err != nil {
				RespError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]any{
				"message": "Data berhasil Disimpan",
				"data":    template.HTML(api.GetData()),
			}
		}

		RespJson(w, http.StatusOK, data)

	}
}

func (api *API) Delete(w http.ResponseWriter, r *http.Request) {

	log.Print("masuk delete")

	r.ParseForm()
	data := make(map[string]any)

	id, _ := strconv.Atoi(r.FormValue("id"))

	log.Println("ini id delete", id)
	err := api.db.Delete(id)
	if err != nil {
		RespError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data = map[string]any{
		"message": "Data berhasil Delete",
		"data":    template.HTML(api.GetData()),
	}

	RespJson(w, http.StatusOK, data)

}

func RespError(w http.ResponseWriter, code int, message string) {
	RespJson(w, code, map[string]string{"error": message})
}

func RespJson(w http.ResponseWriter, code int, send interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(send)
}

func (api *API) IsStructEmpty(s model.Pasien) bool {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.Len() == 0 {
			return true
		}
		if field.Kind() == reflect.Int && field.Int() == 0 {
			return true
		}
		if field.Kind() == reflect.Bool && !field.Bool() {
			return true
		}
	}
	return false
}
