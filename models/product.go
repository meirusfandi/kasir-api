package models

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var Products = []Product{
	{
		ID:    1,
		Name:  "Buku Tulis",
		Price: 5000,
		Stock: 10,
	},
	{
		ID:    2,
		Name:  "Pensil",
		Price: 2000,
		Stock: 20,
	},
	{
		ID:    3,
		Name:  "Penghapus",
		Price: 1000,
		Stock: 30,
	},
}
