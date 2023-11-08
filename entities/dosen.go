package entities

//struct untunk menampung data yang di ambil dari databse

type Dosen struct {
	Id           uint64
	NamaDosen    string
	NIDN         string
	JenisKelamin string
	Alamat       string
	TempatLahir  string
	TanggalLahir string
	NoTlp        string
}
