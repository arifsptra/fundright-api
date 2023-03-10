package campaign

import "website-fundright/user"

// struct for mapping campaign id
type GetCampaignDetailInput struct {
	CampaignID int `uri:"id" binding:"required"`
}

// struct for mapping create campaign
type CreateCampaignInput struct {
	Name string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description string `json:"description" binding:"required"`
	GoalAmount int `json:"goal_amount" binding:"required"`
	Perks string `json:"perks" binding:"required"`
	User user.User
}

// struct for mapping create campaign image
type CreateCampaignImageInput struct {
	CampaignID int `form:"campaign_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User user.User
}