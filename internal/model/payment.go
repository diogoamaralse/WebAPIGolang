package model

type PaymentRequest struct {
	OrderID  string  `json:"orderId"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type PaymentResponse struct {
	ID string `json:"id"`
}
