package models

import (
	"encoding/json"
	"os"
)

type Payment struct {
	Id         int          `json:"id"`
	Amount     float64      `json:"amount"`
	TargetUser UserResponse `json:"user"`
}

type PaymentCollection struct {
	Data []Payment `json:"data"`
}

type PaymentRequest struct {
	Amount   float64 `json:"amount"`
	Username string  `json:"username"`
}

func SavePaymentToJSON(payment Payment) error {
	paymentsCollection, err := GetAllPayments()
	if err != nil {
		return err
	}

	paymentsCollection.Data = append(paymentsCollection.Data, payment)

	data, err := json.MarshalIndent(paymentsCollection, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/payments.json", data, 0644)
}

func GetAllPayments() (PaymentCollection, error) {
	file, err := os.ReadFile("storage/payments.json")
	if err != nil {
		return PaymentCollection{}, err
	}

	var payments PaymentCollection
	err = json.Unmarshal(file, &payments)
	if err != nil {
		return PaymentCollection{}, err
	}

	return payments, nil
}
