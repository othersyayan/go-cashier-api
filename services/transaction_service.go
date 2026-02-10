package services

import (
	"go-cashier-api/models"
	"go-cashier-api/repositories"
)

type TransactionService struct {
	transactionRepo *repositories.TransactionRepository
}

func NewTransactionService(transactionRepo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (s *TransactionService) Checkout(req *models.CheckoutRequest) (*models.Transaction, error) {
	return s.transactionRepo.CreateTransaction(req.Items)
}
