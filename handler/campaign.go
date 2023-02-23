package handler

import (
	"fmt"
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
	// declare input variable
	var input campaign.CreateCampaignInput
	// initialize input data variable with json
	err := c.ShouldBindJSON(&input)
	// error handling
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

	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call function create campaign from service
	newCampaign, err := h.campaignService.CreateCampaign(input)
	// error handling
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

// function to update campaign
func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// declare input id variable
	var inputID campaign.GetCampaignDetailInput
	// initialize input id variable
	err := c.ShouldBindUri(&inputID)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Update Campaign is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// declare input data variable
	var inputData campaign.CreateCampaignInput
	// initialize input data variable with json
	err = c.ShouldBindJSON(&inputData)
	// error handling
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Update Campaign is Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	// call function update data from service
	updateCampaign, err := h.campaignService.UpdateCampaign(inputID, inputData)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Update Campaign is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// response success and send formatter data
	response := helper.APIResponse("Update Campaign is Success!", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))

	// sending json data
	c.JSON(http.StatusOK, response)
}

// function to upload campaign image
func (h *campaignHandler) UploadImage(c *gin.Context) {
	// declare input data
	var input campaign.CreateCampaignImageInput
	// get data from form
	err := c.ShouldBind(&input)
	// error handling
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Campaign Image Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// form file name
	file, err := c.FormFile("campaign_image")
	// error handling
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"is_uploaded": false}
		// response error output
		respons := helper.APIResponse("Campaign Image Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// path file name
	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/campaigns/%d-%s", userID, file.Filename)
	
	// save file
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"is_uploaded": false}
		// response error output
		respons := helper.APIResponse("Campaign Image Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// save campaign image
	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"is_uploaded": false}
		// response error output
		respons := helper.APIResponse("Campaign Image Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}
	
	// formatter error output
	errorMessage := gin.H{"is_uploaded": true}
	// response error output
	respons := helper.APIResponse("Campaign Image Success to Upload!", http.StatusOK, "success", errorMessage)
	c.JSON(http.StatusOK, respons)
}