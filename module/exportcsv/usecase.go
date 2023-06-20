package exportcsv

import (
	"errors"
	"strings"

	"strconv"
)

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) ExportCSV() ([][]string, error) {
	resultGetTransaction, err := u.repo.GetAllTransaction()
	if err != nil {
		return nil, err
	}
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	transactionStringData, err := TransactionStringConverter(resultGetTransaction)
	return transactionStringData, nil
}

func (u UseCase) ExportCSVRangeDateFilter(startDate string, endDate string) ([][]string, error) {
	resultGetTransaction, err := u.repo.GetAllTransactionByRangeDateFilter(startDate, endDate)
	if err != nil {
		return nil, err
	}
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	transactionStringData, err := TransactionStringConverter(resultGetTransaction)
	return transactionStringData, nil

}
func (u UseCase) ExportCSVStatusFilter(status string) ([][]string, error) {
	resultGetTransaction, err := u.repo.GetTransactionByStatusFilter(status)
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	if err != nil {
		return nil, err
	}
	transactionStringData, err := TransactionStringConverter(resultGetTransaction)
	return transactionStringData, nil
}
func (u UseCase) ExportCSVStatusAndRangeDate(status string, startDate string, endDate string) ([][]string, error) {
	resultGetTransaction, err := u.repo.GetTransactionByStatusAndRangeDateFilter(status, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	transactionStringData, err := TransactionStringConverter(resultGetTransaction)
	return transactionStringData, nil

}

func TransactionStringConverter(transactions []Transaction) ([][]string, error) {
	//this define [][] string data and insert header data
	stringData := [][]string{[]string{"id", "od_number", "bank_acount_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "is_hold", "is_fstl_pending", "is_hstl_pending", "is_laa_positif", "payment_amount",
		"billing_gen_date", "is_oda_positif"}}

	//converting [][] data from transaction to be string to be item data
	for _, transaction := range transactions {
		record := []string{strconv.Itoa(transaction.Id), transaction.Oda_number, transaction.Bank_account_no, transaction.Billing_cycle_date,
			strings.Trim(transaction.Payment_due_date, "T00:00:00Z"), strconv.FormatFloat(transaction.Overflow_amount, 'f', 6, 64), strconv.FormatFloat(transaction.Bill_amount, 'f', 6, 64),
			strconv.FormatFloat(transaction.Principal_amount, 'f', 6, 64), strconv.FormatFloat(transaction.Interest_amount, 'f', 6, 64),
			strconv.FormatFloat(transaction.Total_fee_amount, 'f', 6, 64), transaction.Status, transaction.Payment_method, strconv.Itoa(transaction.Auto_debet_counter),
			strings.Trim(transaction.Created_at, "T00:00:00Z"), strings.Trim(transaction.Updated_at, "T00:00:00Z"), strconv.FormatBool(transaction.Is_hold), strconv.FormatBool(transaction.Is_fstl_pending),
			strconv.FormatBool(transaction.Is_hstl_pending), strconv.FormatBool(transaction.Is_laa_positif), strconv.FormatFloat(transaction.Payment_amount, 'f', 6, 64),
			strings.Trim(transaction.Billing_gen_date, "T00:00:00Z"), strconv.FormatBool(transaction.Is_oda_positif)}
		stringData = append(stringData, record)
	}

	return stringData, nil
}
