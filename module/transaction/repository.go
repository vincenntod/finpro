package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	GetAllTransaction() ([]Transaction, error)
	GetTransactionByStatus(status string) ([]Transaction, error)
	GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error)
	GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error)
	GetAllTransactionByDate(start string, end string) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return Repository{db}
}

func (r Repository) GetAllTransaction() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetTransactionByStatus(status string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("status = ?", status).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("created_at BETWEEN ? AND ?", start, end).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Offset((input.Page-1)*input.PageSize).Limit(input.PageSize).Order("created_at desc").Where("status =? AND(created_at BETWEEN ? AND ?)", req.Status, req.StartDate, req.EndDate).Find(&transactions).Error

	return transactions, err
}

func (r Repository) GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Offset((input.Page-1)*input.PageSize).Limit(input.PageSize).Order("created_at desc").Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate).Find(&transactions).Error

	return transactions, err
}
