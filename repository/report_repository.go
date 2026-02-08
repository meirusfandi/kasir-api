package repository

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetReportData(startDate, endDate time.Time) (*models.ReportResponse, error) {
	var report models.ReportResponse

	// Get total revenue and transactions
	queryTotal := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id)
		FROM transactions
		WHERE created_at >= $1 AND created_at <= $2
	`
	err := r.db.QueryRow(queryTotal, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil {
		return nil, err
	}

	// Get best selling product
	queryBestSelling := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as qty_sold
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY p.name
		ORDER BY qty_sold DESC
		LIMIT 1
	`
	err = r.db.QueryRow(queryBestSelling, startDate, endDate).Scan(&report.BestSellingProduct.Name, &report.BestSellingProduct.QtySold)
	if err == sql.ErrNoRows {
		// No transactions or products sold in this period
		report.BestSellingProduct = models.BestSelling{Name: "-", QtySold: 0}
	} else if err != nil {
		return nil, err
	}

	return &report, nil
}
