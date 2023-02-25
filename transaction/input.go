package transaction

import "website-fundright/user"

type CampaignTransaction struct {
	CampaignID int `uri:"id" binding:"required"`
	User       user.User
}

type CreateTransaction struct {
	Amount int `json:"amount" binding:"required"`
	User user.User
	CampaignID int `json:"campaign_id" binding:"required"`
}