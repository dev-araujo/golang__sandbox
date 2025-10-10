package main

import (
	m "fake-poo/models"
	p "fake-poo/processors"
)

func main() {
	paymentData := &m.PaymentData{
		Id:       1,
		Value:    100,
		Currency: "BRL",
	}

	paymentProcessor := &p.PaymentProcessor{}
	paymentProcessor.ProcessPayment(paymentData)

}
