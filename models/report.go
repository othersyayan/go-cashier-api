package models

type SalesReport struct {
	TotalRevenue       int                `json:"total_revenue"`
	TotalTransactions  int                `json:"total_transaksi"`
	BestSellingProduct BestSellingProduct `json:"produk_terlaris"`
}

type BestSellingProduct struct {
	Name     string `json:"nama"`
	Quantity int    `json:"qty_terjual"`
}

type TransactionReport struct {
	TotalRevenue      int           `json:"total_revenue"`
	TotalTransactions int           `json:"total_transaksi"`
	Transactions      []Transaction `json:"transaksi"`
}
