package services

import (
	"kasir-api/models"
	"kasir-api/repository"
	"time"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*models.ReportResponse, error) {
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	return s.repo.GetReportData(startDate, endDate)
}

func (s *ReportService) GetReportByDateRange(startDateStr, endDateStr string) (*models.ReportResponse, error) {
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return nil, err
	}

	// Adjust end date to include the full day
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	return s.repo.GetReportData(startDate, endDate)
}
