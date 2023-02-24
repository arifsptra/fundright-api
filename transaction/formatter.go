package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransaction := CampaignTransactionFormatter{}
	campaignTransaction.ID = transaction.ID
	campaignTransaction.Name = transaction.User.Name
	campaignTransaction.Amount = transaction.Amount
	campaignTransaction.CreatedAt = transaction.CreatedAt
	return campaignTransaction
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	CampaignTransactions := []CampaignTransactionFormatter{}
	for _, transaction := range transactions {
		campaignFormatter := FormatCampaignTransaction(transaction)
		CampaignTransactions = append(CampaignTransactions, campaignFormatter)
	}
	return CampaignTransactions
}