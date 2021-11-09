package dto

import (
	"strings"

	"github.com/andiahmads/bangking-service/helpers"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *helpers.AppError {
	if r.Amount < 5000 {
		return helpers.NewValidationError("To open a new Account your need to deposit atleast 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && r.AccountType != "checking" {
		return helpers.NewValidationError("Account type should be checking or checking")
	}
	return nil
}
