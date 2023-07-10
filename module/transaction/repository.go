package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	GetAllTransaction() ([]Transaction, error)
	GetAllTransactionByStatus(status string) ([]Transaction, error)
	GetAllTransactionByDate(start string, end string) ([]Transaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error)
	GetAllLimit(input FilterLimit) ([]Transaction, error, int64)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return Repository{db}
}

func (r Repository) GetAllTransaction() ([]Transaction, error) {
	var transactions []Transaction

	err := r.DB.Find(&transactions).Error

	return transactions, err
}

func (r Repository) GetAllTransactionByStatus(status string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.DB.Where("status = ?", status).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.DB.Where("status =? AND(created_at BETWEEN ? AND ?)", status, start, end).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllLimit(input FilterLimit) ([]Transaction, error, int64) {
	var transactions []Transaction
	var count int64

	r.DB.Model(&transactions).Count(&count)
	err := r.DB.Offset((input.Page - 1) * input.PageSize).Limit(input.PageSize).Order("created_at desc").Find(&transactions).Error
	return transactions, err, count
}
