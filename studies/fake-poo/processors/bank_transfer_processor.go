package processors

import (
	"fake-poo/models"
	"fmt"
)

type BankTransferProcessor struct{}

func (p *BankTransferProcessor) ProcessPayment(data *models.PaymentData) (string, error) {
	fmt.Printf("Processando pagamento via Transferência Bancária... ID: %d\n", data.Id)
	return "bt-transaction-456", nil
}
