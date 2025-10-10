package processors

import (
	m "fake-poo/models"
	"fmt"
)

type PaymentProcessor struct{}

func (pp *PaymentProcessor) ProcessPayment(data *m.PaymentData) (string, error) {

	fmt.Printf("Processando pagamento... ID: %d, Valor: %d, Moeda: %s\n", data.Id, data.Value, data.Currency)

	return "transacao-abc-123", nil
}
