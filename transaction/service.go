package transaction

type Service interface {
	GetTransactionByCampaignID(input CampaignTransaction) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignID(input CampaignTransaction) ([]Transaction, error) {
	transactions, err := s.repository.FindByCampaignID(input.CampaignID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
