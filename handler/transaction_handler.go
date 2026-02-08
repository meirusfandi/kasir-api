package handler

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// multiple item apa aja, quantity nya
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		helpers.SendResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	transaction, err := h.service.Checkout(req.Items, false)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Checkout successful", transaction)
}
