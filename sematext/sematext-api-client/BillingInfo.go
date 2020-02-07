package sematext

type BillingInfo struct {
	CreditCardId  int64  `json:"creditCardId"`
	PaymentMethod string `json:"paymentMethod"`
	PlanId        int64  `json:"planId"`
}
