package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) FindAll() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r Repository) FindByStatus(status string) (Transaction, error) {
	var transactions Transaction
	err := r.db.Where("status = ?", status).Find(&transactions).Error
	return transactions, err
}
