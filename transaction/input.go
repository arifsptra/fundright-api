package transaction

import "website-fundright/user"

type CampaignTransaction struct {
	CampaignID int `uri:"id" binding:"required"`
	User       user.User
}