package transaction

import (
	"golang/module/transaction/entities"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	GetAllTransaction(page int, size int) ([]entities.Transaction, error)
	GetAllTransactionByStatus(status string, page int, size int) ([]entities.Transaction, error)
	GetAllTransactionByDate(start string, end string, page int, size int) ([]entities.Transaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) ([]entities.Transaction, error)

	GetAllTransactionNoLimit() ([]entities.Transaction, error)
	GetAllTransactionByStatusNoLimit(status string) ([]entities.Transaction, error)
	GetAllTransactionByDateNoLimit(start string, end string) ([]entities.Transaction, error)
	GetAllTransactionByStatusDateNoLimit(status string, start string, end string) ([]entities.Transaction, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return Repository{db}
}

func (r Repository) GetAllTransaction(page int, size int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Offset((page - 1) * size).Limit(size).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByStatus(status string, page int, size int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Where("status = ?", status).Offset((page - 1) * size).Limit(size).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByDate(start string, end string, page int, size int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Where("created_at BETWEEN ? AND ?", start, end).Offset((page - 1) * size).Limit(size).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Where("status =? AND(created_at BETWEEN ? AND ?)", status, start, end).Offset((page - 1) * size).Limit(size).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionNoLimit() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByStatusNoLimit(status string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Where("status = ?", status).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByDateNoLimit(start string, end string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByStatusDateNoLimit(status string, start string, end string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.DB.Where("status =? AND(created_at BETWEEN ? AND ?)", status, start, end).Find(&transactions).Error
	return transactions, err
}
