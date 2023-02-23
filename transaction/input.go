package transaction

type CampaignTransaction struct {
	CampaignID int `uri:"id" binding:"required"`
}