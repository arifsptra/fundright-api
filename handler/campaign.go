package handler

import (
	"net/http"
	"strconv"
	"website-fundright/campaign"
	"website-fundright/helper"

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
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// query to get user_id
	userID, _ := strconv.Atoi(c.Query("user_id"))

	// call funtion to get campaign from campaign service
	campaigns, err := h.campaignService.GetCampaign(userID)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Get Campaigns is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// response success
	response := helper.APIResponse("Get Campaigns is Success!", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))

	// sending json data
	c.JSON(http.StatusOK, response)
}