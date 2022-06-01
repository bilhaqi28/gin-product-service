package request

import "mime/multipart"

type UpdateProduct struct {
	BarangKode      string `form:"barang_kode" json:"barang_kode" binding:"required"`
	BarangNama      string `form:"barang_nama" json:"barang_nama"`
	BarangDesc      string `form:"barang_desc" json:"barang_desc"`
	BarangThumbnail string
	BarangFoto      string                `form:"barang_foto" json:"barang_foto"`
	Thumbnail       *multipart.FileHeader `form:"thumbnail" json:"thumbnail"`
	BarangId        int
}
