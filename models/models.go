package models

type Barang struct {
	IdBarang   uint   `json:"idBarang" gorm:"primary_key"`
	NamaBarang string `json:"namaBarang"`
	Satuan     string `json:"satuan"`
	Harga      uint   `json:"harga"`
	Items      []Item `json:"items" gorm:"foreignkey:Id_Barang"`
}
type Item struct {
	ItemId      uint   `json:"itemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	IdBarang    uint   `json:"-"`
}
