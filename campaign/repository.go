package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByCampaignID(campaignID int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function to find all the campaign
func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// function to find campaign by user id
func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

// function to find campaign by campaign id
func (r *repository) FindByCampaignID(campaignID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", campaignID).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// function to create campaign
func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}