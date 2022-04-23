package response

type ProductWeb struct {
	Id              int    `json:"id"`
	BarangKode      string `json:"barang_kode"`
	BarangNama      string `json:"barang_nama"`
	BarangDesc      string `json:"barang_desc"`
	BarangThumbnail string `json:"barang_thumbnail"`
	BarangFoto      string `json:"barang_foto"`
}
