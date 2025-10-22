package processors

import (
	"fake-poo/models"
	"fmt"
)

type PixProcessor struct{}

func (p *PixProcessor) ProcessPayment(data *models.PaymentData) (string, error) {
	fmt.Printf("Processando pagamento via Pix... ID: %d\n", data.Id)
	return "cc-transaction-123", nil
}
