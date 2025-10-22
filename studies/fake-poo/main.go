package main

import (
	m "fake-poo/models"
	p "fake-poo/processors"
	"fmt"
)

func main() {
	paymentData := &m.PaymentData{
		Id:       1,
		Value:    100,
		Currency: "BRL",
	}
	creditCard := &p.CreditCardProcessor{}
	bankTransfer := &p.BankTransferProcessor{}

	transactionOperation(creditCard, paymentData)
	transactionOperation(bankTransfer, paymentData)

}

func transactionOperation(method p.PaymentMethod, data *m.PaymentData) {
	transactionID, err := method.ProcessPayment(data)
	if err != nil {
		fmt.Println("Erro ao processar:", err)
		return
	}
	fmt.Printf("Transação concluída com sucesso! ID: %s\n", transactionID)
}
