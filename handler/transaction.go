package handler

import (
	"net/http"
	"website-fundright/helper"
	"website-fundright/transaction"
	"website-fundright/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

// function to get campaign transaction by campaign id
func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	// declare input data
	var input transaction.CampaignTransaction

	// initialize input id variable
	err := c.ShouldBindUri(&input)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get Campaign Transaction is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call function to get campaign transaction by campaign id in service
	transactions, err := h.service.GetTransactionByCampaignID(input)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get Campaign Transaction is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// response error output
	respons := helper.APIResponse("Campaign Transaction Success to Get!", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, respons)
}