package processors

import (
	m "fake-poo/models"
)

type PaymentMethod interface {
	ProcessPayment(data *m.PaymentData) (string, error)
}
