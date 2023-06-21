package exportcsv

import (
	"errors"
	"strconv"
	"strings"
)

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) ExportCSV(req *ExportCSVRequest) ([][]string, error) {

	var resultGetTransaction []Transaction
	var transactionStringData [][]string
	var err error
	switch {
	case req.Status == "" && req.StartDate == "" && req.EndDate == "": //without filter
		resultGetTransaction, err = u.repo.GetAllTransaction()
	case req.Status == "": // rage date filter
		resultGetTransaction, err = u.repo.GetAllTransactionByRangeDateFilter(req)
	case req.StartDate == "" && req.EndDate == "": //only status filter
		resultGetTransaction, err = u.repo.GetTransactionByStatusFilter(req)
	default: //complate filter(status, rage date)
		resultGetTransaction, err = u.repo.GetTransactionByStatusAndRangeDateFilter(req)
	}
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	transactionStringData, err = TransactionStringConverter(resultGetTransaction)
	if err != nil {
		return nil, err
	}
	return transactionStringData, nil
}

func TransactionStringConverter(transactions []Transaction) ([][]string, error) {
	//set name field transaction in first record
	stringData := [][]string{[]string{"id", "od_number", "bank_acount_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "is_hold", "is_fstl_pending", "is_hstl_pending", "is_laa_positif", "payment_amount",
		"billing_gen_date", "is_oda_positif"}}
	//converting some data from transaction to string
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
