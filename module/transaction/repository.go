package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	GetAllTransaction() ([]Transaction, error)
<<<<<<< HEAD
	GetAllTransactionByStatus(req *FilterByStatusDate) ([]Transaction, error)
	GetAllTransactionByDate(req *FilterByStatusDate) ([]Transaction, error)
	GetAllTransactionByStatusDate(req *FilterByStatusDate) ([]Transaction, error)
=======
	GetTransactionByStatus(status string) ([]Transaction, error)
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
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

<<<<<<< HEAD
func (r Repository) GetAllTransactionByStatus(req *FilterByStatusDate) ([]Transaction, error) {
=======
func (r Repository) GetTransactionByStatus(status string) ([]Transaction, error) {
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
	var transactions []Transaction
	err := r.db.Where("status = ?", req.Status).Find(&transactions).Error
	return transactions, err
}

func (r Repository) GetAllTransactionByDate(req *FilterByStatusDate) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate).Find(&transactions).Error
	return transactions, err
}

<<<<<<< HEAD
func (r Repository) GetAllTransactionByStatusDate(req *FilterByStatusDate) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("status =? AND(created_at BETWEEN ? AND ?)", req.Status, req.StartDate, req.EndDate).Find(&transactions).Error
	return transactions, err
}


=======
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
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
