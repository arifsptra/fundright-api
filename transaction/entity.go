package transaction

import (
	"website-fundright/user"
	"time"
)

type Transaction struct {
	ID         int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     int
	CampaignID int
	User user.User
}