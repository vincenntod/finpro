package exportcsv

import (
	"fmt"
	"golang/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetAllTransaction() ([]Transaction, error) {
	var transactions []Transaction
	if err := model.DB.Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return nil, err

	}
	//	fmt.Println(transactions)
	return transactions, nil
}
func (r Repository) GetTransactionByStatusFilter(status string) ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Where("status = ?", status).Find(&transactions).Error; err != nil {
		return nil, err

	}

	return transactions, nil
}
func (r Repository) GetAllTransactionByRangeDateFilter(startDate string, endDate string) ([]Transaction, error) {
	fmt.Println(startDate, endDate)
	var transactions []Transaction
	if err := model.DB.Where("created_at BETWEEN  ? AND  ?", startDate, endDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r Repository) GetTransactionByStatusAndRangeDateFilter(status string, startDate string, endDate string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("status =? AND(created_at  BETWEEN ? AND ?)", status, startDate, endDate).Find(&transactions).Error
	if err != nil {
		return nil, err

	}

	return transactions, nil
}
