package processors

import (
	"fake-poo/models"
	"fmt"
)

type CreditCardProcessor struct{}

func (p *CreditCardProcessor) ProcessPayment(data *models.PaymentData) (string, error) {
	fmt.Printf("Processando pagamento via Cartão de Crédito... ID: %d\n", data.Id)
	return "cc-transaction-123", nil
}
