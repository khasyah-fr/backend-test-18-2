package models

import "time"

type Pelanggan struct {
	ID            uint `gorm:"primaryKey"`
	NamaPelanggan string
	NomorTelepon  string
	Alamat        string
}

type Pesanan struct {
	ID             uint `gorm:"primaryKey"`
	PelangganID    uint
	TanggalPesanan time.Time
	TotalHarga     float64
	DetailPesanan  []DetailPesanan `gorm:"foreignKey:PesananID"`
	StrukPembelian StrukPembelian
}

type Menu struct {
	ID         uint `gorm:"primaryKey"`
	NamaMenu   string
	HargaMenu  float64
	JumlahStok int
	Pesanan    []Pesanan `gorm:"many2many:detail_pesanan"`
}

type DetailPesanan struct {
	ID            uint `gorm:"primaryKey"`
	PesananID     uint
	MenuID        uint
	JumlahPesanan int
}

type StrukPembelian struct {
	ID               uint `gorm:"primaryKey"`
	PesananID        uint
	TanggalPembelian time.Time
	TotalHarga       float64
}

type LaporanPenghasilan struct {
	ID               uint `gorm:"primaryKey"`
	TanggalLaporan   time.Time
	TotalPenghasilan float64
	Pesanan          []Pesanan `gorm:"many2many:detail_laporan_penghasilan"`
}

type LaporanStok struct {
	ID             uint `gorm:"primaryKey"`
	TanggalLaporan time.Time
	JumlahStok     int
	Menu           []Menu `gorm:"many2many:detail_laporan_stok"`
}
