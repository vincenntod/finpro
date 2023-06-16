package exportcsv

import (
	"errors"
	"fmt"
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

func (u UseCase) ExportCSV(req *ExportCSVRequest) ([][]string, error) {
	fmt.Println("case", req)
	exportData := [][]string{[]string{"id", "od_number", "bank_acount_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "is_hold", "is_fstl_pending", "is_hstl_pending", "is_laa_positif", "payment_amount",
		"billing_gen_date", "is_oda_positif"}}
	var resultGetTransaction []Transaction
	switch {
	case req.Status == "" && req.StartDate == "" && req.EndDate == "":
		data, err := u.repo.GetAllTransaction()
		if err != nil {
			return nil, err
		}
		resultGetTransaction = data

	case req.Status == "":
		data, err := u.repo.GetAllTransactionByStartAndEndDate(req.StartDate, req.EndDate)
		if err != nil {
			return nil, err
		}
		resultGetTransaction = data

	case req.StartDate == "" && req.EndDate == "":
		data, err := u.repo.GetTransactionByStatus(req)
		if err != nil {
			return nil, err
		}
		resultGetTransaction = data
	default:
		data, err := u.repo.GetTransactionByStatusAndStartAndEndDate(req.Status, req.StartDate, req.EndDate)
		if err != nil {
			return nil, err
		}

		resultGetTransaction = data
	}
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	for _, transaction := range resultGetTransaction {
		record := []string{strconv.Itoa(transaction.Id), transaction.Oda_number, transaction.Bank_account_no, transaction.Billing_cycle_date,
			transaction.Payment_due_date.GoString(), strconv.FormatFloat(transaction.Overflow_amount, 'f', 6, 64), strconv.FormatFloat(transaction.Bill_amount, 'f', 6, 64),
			strconv.FormatFloat(transaction.Principal_amount, 'f', 6, 64), strconv.FormatFloat(transaction.Interest_amount, 'f', 6, 64),
			strconv.FormatFloat(transaction.Total_fee_amount, 'f', 6, 64), transaction.Status, transaction.Payment_method, strconv.Itoa(transaction.Auto_debet_counter),
			transaction.Created_at.GoString(), transaction.Updated_at.GoString(), strconv.FormatBool(transaction.Is_hold), strconv.FormatBool(transaction.Is_fstl_pending),
			strconv.FormatBool(transaction.Is_hstl_pending), strconv.FormatBool(transaction.Is_laa_positif), strconv.FormatFloat(transaction.Payment_amount, 'f', 6, 64),
			transaction.Billing_gen_date.GoString(), strconv.FormatBool(transaction.Is_oda_positif)}
		exportData = append(exportData, record)
	}
	return exportData, nil

}
