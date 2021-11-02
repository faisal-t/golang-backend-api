package handler

import (
	"bwa-golang/helper"
	"bwa-golang/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to get Transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transaction, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("transaction details", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)
}
