package utils

import (
	"github.com/diegocabrera89/ms-payment-bank-sandbox/constantsmicro"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/models"
	"github.com/google/uuid"
	"time"
)

// BuildBank build bank object.
func BuildBank(bank *models.Bank) {
	bank.BankID = uuid.New().String()        // Generate a unique ID for the bank
	bank.CreatedAt = time.Now().UTC().Unix() //Date in UTC
	bank.Status = constantsmicro.StatusBankEnable
}
