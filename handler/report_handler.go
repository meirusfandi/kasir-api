package handler

import (
	"kasir-api/helpers"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetDailyReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetDailyReport()
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Get daily report", report)
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		helpers.SendResponse(w, http.StatusBadRequest, "start_date and end_date are required", nil)
		return
	}

	report, err := h.service.GetReportByDateRange(startDate, endDate)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Get report by date range", report)
}
