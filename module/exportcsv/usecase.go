package exportcsv

import (
	"fmt"
	"strconv"
)

// dumy data
type TransactionItem struct {
	Id         int
	Oda_number string
	Status     string
	Price      float64
	Total_data int
	Created_at string
}

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) ExportCSV(req *ExportCSVRequest) ([][]string, error) {
	fmt.Println("case", req)
	exportData := [][]string{[]string{"id", "oda_number", "status", "price", "total_data", "created_at"}}
	var resultGetTransaction []Transaction
	switch req.Status {
	case "":
		fmt.Println("req", req.Status, "!")
		result, err := u.repo.GetAllTransaction()
		if err != nil {
			return nil, err
		}
		resultGetTransaction = result
	case "success", "pending", "reject":
		result, err := u.repo.GetTransactionByStatus(req)
		if err != nil {
			return nil, err

		}
		resultGetTransaction = result
	}

	for _, transaction := range resultGetTransaction {
		record := []string{strconv.Itoa(transaction.Id), transaction.Oda_number, transaction.Status, strconv.Itoa(int(transaction.Price)), strconv.Itoa(transaction.Total_data), transaction.Created_at}
		exportData = append(exportData, record)
	}
	return exportData, nil

}
