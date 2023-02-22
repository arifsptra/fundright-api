package handler

import (
	"net/http"
	"strconv"
	"website-fundright/campaign"
	"website-fundright/helper"
	"website-fundright/user"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

// pass this struct as a service parameter
func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

// function to get campaign
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// query to get user_id
	userID, _ := strconv.Atoi(c.Query("user_id"))

	// call funtion to get campaign from campaign service
	campaigns, err := h.campaignService.GetCampaigns(userID)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get Campaigns is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// response success
	response := helper.APIResponse("Get Campaigns is Success!", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))

	// sending json data
	c.JSON(http.StatusOK, response)
}

// function to get campaign by campaign id
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// declare input variable
	var input campaign.GetCampaignDetailInput
	// initialize input variable
	err := c.ShouldBindUri(&input)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get Campaign is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// get campaign by campaign id
	campaignDetail, err := h.campaignService.GetCampaignByID(input)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get Campaign is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// response success
	response := helper.APIResponse("Get Campaign is Success!", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))

	// sending json data
	c.JSON(http.StatusOK, response)
}

// function to create new campaign
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Create Campaign is Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		// response error output
		respons := helper.APIResponse("Create Campaign is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// response success
	response := helper.APIResponse("Create Campaign is Success!", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))

	// sending json data
	c.JSON(http.StatusOK, response)
}