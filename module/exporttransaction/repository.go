package exporttransaction

import (
	"fmt"
	"golang/module/exporttransaction/dto"
	"golang/module/exporttransaction/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetAllTransaction() ([]entity.Transaction, error)
	GetTransactionByStatusFilter(req *dto.ExportCSVRequest) ([]entity.Transaction, error)
	GetAllTransactionByRangeDateFilter(req *dto.ExportCSVRequest) ([]entity.Transaction, error)
	GetTransactionByStatusAndRangeDateFilter(req *dto.ExportCSVRequest) ([]entity.Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetAllTransaction() ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return transactions, nil
}

func (r repository) GetTransactionByStatusFilter(req *dto.ExportCSVRequest) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := r.db.Where("status = ?", req.Filter.Status).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r repository) GetAllTransactionByRangeDateFilter(req *dto.ExportCSVRequest) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := r.db.Where("created_at BETWEEN  ? AND  ?", req.Filter.StartDate, req.Filter.EndDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r repository) GetTransactionByStatusAndRangeDateFilter(req *dto.ExportCSVRequest) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Where("status =? AND(created_at  BETWEEN ? AND ?)", req.Filter.Status, req.Filter.StartDate, req.Filter.EndDate).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
