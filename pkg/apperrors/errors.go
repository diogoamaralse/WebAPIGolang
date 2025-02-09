package apperrors

import "errors"

var (
	ErrInvalidPaymentRequest = errors.New("invalid payment request")
	ErrPaymentProcessing     = errors.New("error processing payment")
)
