package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/yudiagonal/go-crud-form-validation/entities"
	"github.com/yudiagonal/go-crud-form-validation/models"
)

var dosenModels = models.NewDosenModels()

func Index(w http.ResponseWriter, r *http.Request) {
	dosen, _ := dosenModels.FindAll()

	data := map[string]interface{}{
		"dosen": dosen,
	}

	tmpl, err := template.ParseFiles("views/dosen/index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)

}
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("views/dosen/create.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var dosen entities.Dosen
		dosen.NamaDosen = r.Form.Get("nama_dosen")
		dosen.NIDN = r.Form.Get("nidn")
		dosen.JenisKelamin = r.Form.Get("jenis_kelamin")
		dosen.Alamat = r.Form.Get("alamat")
		dosen.TempatLahir = r.Form.Get("tempat_lahir")
		dosen.TanggalLahir = r.Form.Get("tanggal_lahir")
		dosen.NoTlp = r.Form.Get("no_tlp")
		fmt.Println(dosen)

		dosenModels.Create(dosen)
		data := map[string]interface{}{
			"pesan": "Data Dosen berhasil disimpan",
		}

		tmpl, _ := template.ParseFiles("views/dosen/create.html")
		tmpl.Execute(w, data)

	}

}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
