package repositories

import (
	"database/sql"

	"go-cashier-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetSalesReport(startDate, endDate string) (*models.SalesReport, error) {
	report := &models.SalesReport{}

	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0), 
			COUNT(id) 
		FROM transactions 
		WHERE created_at::date BETWEEN $1 AND $2`

	err := repo.db.QueryRow(query, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil {
		return nil, err
	}

	productQuery := `
		SELECT 
			p.name, 
			SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at::date BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1`

	err = repo.db.QueryRow(productQuery, startDate, endDate).Scan(&report.BestSellingProduct.Name, &report.BestSellingProduct.Quantity)
	if err == sql.ErrNoRows {
		report.BestSellingProduct.Name = "-"
		report.BestSellingProduct.Quantity = 0
	} else if err != nil {
		return nil, err
	}

	return report, nil
}
