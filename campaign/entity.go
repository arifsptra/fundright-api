package campaign

import (
	// "os/user"
	"time"
	"website-fundright/user"
)

// struct for Campaign data
type Campaign struct {
	ID int
	Name string
	ShortDescription string
	Desciption string
	GoalAmount int
	CurrentAmount int
	Perks string
	BackerCount int
	Slug string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID int
	CampaignImages []CampaignImage
	User user.User
}

// struct for Campaign Image data
type CampaignImage struct {
	ID int
	FileName string
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
	CampaignID int
}