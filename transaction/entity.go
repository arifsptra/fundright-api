package transaction

import (
	"time"
	"website-fundright/campaign"
	"website-fundright/user"
)

type Transaction struct {
	ID         int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     int
	CampaignID int
	User user.User
	Campaign campaign.Campaign
}