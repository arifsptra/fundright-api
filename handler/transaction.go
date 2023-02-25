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

// function to get user transction by user id
func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// call function to get user transaction by user is in service
	transactions, err := h.service.GetTransactionByUserID(userID)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get User Transaction is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// response error output
	respons := helper.APIResponse("User Transaction Success to Get!", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, respons)
}

// function to create transaction
func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	// desclare transaction input data
	var input transaction.CreateTransaction

	// iinitiate transaction
	err := c.ShouldBindJSON(&input)
	// error handling
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Create Transaction is Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call function to get user transaction by user is in service
	newTransaction, err := h.service.CreateTransaction(input)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Create Transaction is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// response error output
	respons := helper.APIResponse("Create Transaction is Success!", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, respons)
}