package services

import (
	"time"

	"go-cashier-api/models"
	"go-cashier-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport() (*models.SalesReport, error) {
	today := time.Now().Format("2006-01-02")
	return s.repo.GetSalesReport(today, today)
}

func (s *ReportService) GetReportByRange(startDate, endDate string) (*models.SalesReport, error) {
	return s.repo.GetSalesReport(startDate, endDate)
}
