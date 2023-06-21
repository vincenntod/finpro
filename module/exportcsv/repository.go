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
func (r Repository) GetTransactionByStatusFilter(req *ExportCSVRequest) ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Where("status = ?", req.Status).Find(&transactions).Error; err != nil {
		return nil, err

	}

	return transactions, nil
}
func (r Repository) GetAllTransactionByRangeDateFilter(req *ExportCSVRequest) ([]Transaction, error) {
	var transactions []Transaction
	if err := model.DB.Where("created_at BETWEEN  ? AND  ?", req.StartDate, req.EndDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r Repository) GetTransactionByStatusAndRangeDateFilter(req *ExportCSVRequest) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("status =? AND(created_at  BETWEEN ? AND ?)", req.Status, req.StartDate, req.EndDate).Find(&transactions).Error
	if err != nil {
		return nil, err

	}

	return transactions, nil
}
