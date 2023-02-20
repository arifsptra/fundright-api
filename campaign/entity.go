package campaign

import "time"

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