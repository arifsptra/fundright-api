package campaign

// struct for mapping campaign id
type GetCampaignDetailInput struct {
	CampaignID int `uri:"id" binding:"required"`
}