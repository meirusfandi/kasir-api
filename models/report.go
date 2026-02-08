package models

type BestSelling struct {
	Name    string `json:"nama"`
	QtySold int    `json:"qty_terjual"`
}

type ReportResponse struct {
	TotalRevenue       int         `json:"total_revenue"`
	TotalTransactions  int         `json:"total_transaksi"`
	BestSellingProduct BestSelling `json:"produk_terlaris"`
}
