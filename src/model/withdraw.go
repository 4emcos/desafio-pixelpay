package model

type Withdraw struct {
	Value float64 `json:"value" validate:"required, min=1"`
	Payee string  `json:"payee" validate:"required, min=11, max=11"` // beneficiary
}
