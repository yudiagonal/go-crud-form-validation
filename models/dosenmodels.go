package models

import (
	"database/sql"
	"fmt"

	"github.com/yudiagonal/go-crud-form-validation/config"
	"github.com/yudiagonal/go-crud-form-validation/entities"
)

//untuk menyimpan inputan dari entiti dosen ke database

type DosenModels struct {
	conn *sql.DB
}

func NewDosenModels() *DosenModels {
	conn, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	return &DosenModels{
		conn: conn,
	}
}

// method untuk mengambil data pasien
func (d *DosenModels) FindAll() ([]entities.Dosen, error) {

	rows, err := d.conn.Query("select * from dosen")
	if err != nil {
		return []entities.Dosen{}, err
	}
	defer rows.Close()

	var dataDosen []entities.Dosen
	for rows.Next() {
		var dosen entities.Dosen
		rows.Scan(
			&dosen.Id,
			&dosen.NamaDosen,
			&dosen.NIDN,
			&dosen.JenisKelamin,
			&dosen.Alamat,
			&dosen.TempatLahir,
			&dosen.TanggalLahir,
			&dosen.NoTlp,
		)
		dataDosen = append(dataDosen, dosen)
	}
	return dataDosen, nil
}

func (d *DosenModels) Create(dosen entities.Dosen) bool {
	result, err := d.conn.Exec("INSERT INTO dosen(nama_dosen, nidn, jenis_kelamin, alamat, tempat_lahir, tanggal_lahir, no_tlp) values(?,?,?,?,?,?,?)", dosen.NamaDosen, dosen.NIDN, dosen.JenisKelamin, dosen.Alamat, dosen.TempatLahir, dosen.TanggalLahir, dosen.NoTlp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}
