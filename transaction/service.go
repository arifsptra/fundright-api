package transaction

import (
	"errors"
	"website-fundright/campaign"
)

type Service interface {
	GetTransactionByCampaignID(input CampaignTransaction) ([]Transaction, error)
}

type service struct {
	repository Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input CampaignTransaction) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByCampaignID(input.CampaignID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}
	
	transactions, err := s.repository.FindByCampaignID(input.CampaignID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
