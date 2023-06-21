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
<<<<<<< HEAD
func (r Repository) GetTransactionByStatusFilter(status string) ([]Transaction, error) {
=======
func (r Repository) GetTransactionByStatus(req *ExportCSVRequest) ([]Transaction, error) {
>>>>>>> parent of 007e65b (solve error and adding custom resspone invalid json format)
	var transactions []Transaction
	if err := r.db.Where("status = ?", status).Find(&transactions).Error; err != nil {
		return nil, err

	}

	return transactions, nil
}
func (r Repository) GetAllTransactionByStartAndEndDate(startDate string, endDate string) ([]Transaction, error) {
	fmt.Println(startDate, endDate)
	var transactions []Transaction
	if err := model.DB.Where("created_at BETWEEN  ? AND  ?", startDate, endDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r Repository) GetTransactionByStatusAndStartAndEndDate(status string, startDate string, endDate string) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("status =? AND(created_at  BETWEEN ? AND ?)", status, startDate, endDate).Find(&transactions).Error
	if err != nil {
		return nil, err

	}

	return transactions, nil
}
