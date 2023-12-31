package model

type Transaction struct {
	Value float64 `json:"value" validate:"required, min=1"`
	Payer string  `json:"payer" validate:"required, min=11, max=11"` // payer
	Payee string  `json:"payee" validate:"required, min=11, max=11"` // beneficiary
}
