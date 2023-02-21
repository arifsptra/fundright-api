package campaign

type Service interface {
	GetCampaign(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// function to get campaign
func (s *service) GetCampaign(userID int) ([]Campaign, error) {
	// if userID is not 0 call the function FindByID
	if userID != 0 {
		campaigns, err := s.repository.FindByID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	// if userID is 0 call the function FindAll
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}