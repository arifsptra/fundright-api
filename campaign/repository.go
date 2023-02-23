package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByCampaignID(campaignID int) (Campaign, error)
	CreateCampaignImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(CampaignID int) (bool, error)
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

// function to update campaign
func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// function to create/upload campaign images
func (r *repository) CreateCampaignImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

// function to change is_primary to false
func (r *repository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	// UPDATE campaign_images SET is_primary=false WHERE campaign_id=1
	err := r.db.Model(&CampaignImage{}).Where("campaign_id=?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
