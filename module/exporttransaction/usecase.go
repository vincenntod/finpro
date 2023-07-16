package exporttransaction

import (
	"errors"

	"golang/module/exporttransaction/dto"
	"golang/module/exporttransaction/entity"

	"strconv"
	"strings"
)

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *useCase {
	return &useCase{
		repo: repo,
	}
}

type Usecase interface {
	ExportCSV(req *dto.ExportCSVRequest) ([][]string, error)
}

func (u useCase) ExportCSV(req *dto.ExportCSVRequest) ([][]string, error) {

	var resultGetTransaction []entity.Transaction
	var transactionStringData [][]string
	var err error

	switch {
	case req.Filter.Status == "" && req.Filter.StartDate == "" && req.Filter.EndDate == "":
		resultGetTransaction, err = u.repo.GetAllTransaction()
	case req.Filter.Status == "":
		resultGetTransaction, err = u.repo.GetAllTransactionByRangeDateFilter(req)
	case req.Filter.StartDate == "" && req.Filter.EndDate == "":
		resultGetTransaction, err = u.repo.GetTransactionByStatusFilter(req)
	default:
		resultGetTransaction, err = u.repo.GetTransactionByStatusAndRangeDateFilter(req)
	}
	if err != nil {
		return nil, err
	}

	if len(resultGetTransaction) == 0 {
		return nil, errors.New("not found")
	}
	transactionStringData, err = TransactionStringConverter(resultGetTransaction)
	if err != nil {
		return nil, err
	}
	return transactionStringData, nil
}

func TransactionStringConverter(transactions []entity.Transaction) ([][]string, error) {
	stringData := [][]string{{"id", "oda_number", "status", "price", "created_at"}}
	for _, transaction := range transactions {
		record := []string{strconv.Itoa(transaction.Id), transaction.Oda_number, transaction.Status, strconv.FormatFloat(transaction.Bill_amount, 'f', 6, 64), strings.Trim(transaction.Created_at, "T00:00:00Z")}
		stringData = append(stringData, record)
	}
	return stringData, nil
}
