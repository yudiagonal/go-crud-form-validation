package entities

//struct untunk menampung data yang di ambil dari databse

type Dosen struct {
	Id           uint64
	NamaDosen    string
	Nidn         string
	JenisKelamin string
	Alamant      string
	TempatLahir  string
	TanggalLahir string
	NoTlp        string
}
