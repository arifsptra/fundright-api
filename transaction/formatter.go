package transaction

import (
	"time"
)

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

type CampaignUserTransaction struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

type UserTransactionFormatter struct {
	ID int `json:"id"`
	Amount int `json:"amount"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Campaign CampaignUserTransaction `json:"campaign"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	userTransaction := UserTransactionFormatter{}
	userTransaction.ID = transaction.ID
	userTransaction.Amount = transaction.Amount
	userTransaction.Status = transaction.Status
	userTransaction.CreatedAt = transaction.CreatedAt
	userTransaction.Campaign.Name = transaction.Campaign.Name
	userTransaction.Campaign.ImageURL = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		userTransaction.Campaign.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}
	return userTransaction
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	userTransactions := []UserTransactionFormatter{}
	for _, transaction := range transactions {
		userTransaction := FormatUserTransaction(transaction)
		userTransactions = append(userTransactions, userTransaction)
	}
	return userTransactions
}

type TransactionFormatter struct {
	ID int `json:"id"`
	CampaignID int `json:"campaign_id"`
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
	Status string `json:"status"`
	Code string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL
	return formatter
}