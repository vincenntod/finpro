package transaction

import (
	"strconv"
	"strings"
)

type UseCase struct {
	repo RepositoryInterface
}

type UseCaseInterface interface {
<<<<<<< HEAD
	GetAllTransactionOnePoint(req *FilterByStatusDate) ([]Transaction, error)
=======
	GetAllTransaction() ([]Transaction, error)
	GetTransactionByStatus(status string) ([]Transaction, error)
	GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error)
	GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error)
	GetAllTransactionByDate(start string, end string) ([]Transaction, error)
<<<<<<< HEAD
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
=======
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
}

func NewUseCase(repo RepositoryInterface) UseCaseInterface {
	return UseCase{
		repo: repo,
	}
}

func (u UseCase) GetAllTransactionOnePoint(req *FilterByStatusDate) ([]Transaction, error) {
	switch {
	case req.Status == "" && req.StartDate == "" && req.EndDate == "":
		return u.repo.GetAllTransaction()
	case req.Status == "":
		return u.repo.GetAllTransactionByDate(req)
	case req.StartDate == "" && req.EndDate == "":
		return u.repo.GetAllTransactionByStatus(req)
	default:
		return u.repo.GetAllTransactionByStatusDate(req)
	}
	
}

<<<<<<< HEAD
<<<<<<< HEAD
func TransactionStringConverter(transactions []Transaction) ([][]string, error) {
	//set name field transaction in first record
	stringData := [][]string{[]string{"id", "od_number", "bank_acount_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "is_hold", "is_fstl_pending", "is_hstl_pending", "is_laa_positif", "payment_amount",
		"billing_gen_date", "is_oda_positif"}}
	//converting some data from transaction to string
	for _, transaction := range transactions {
		record := []string{strconv.Itoa(transaction.Id), transaction.OdaNumber, transaction.BankAccountNo, transaction.BillingCycleDate,
			strings.Trim(transaction.PaymentDueDate, "T00:00:00Z"), strconv.FormatFloat(float64(transaction.OverflowAmount), 'f', 6, 64), strconv.FormatFloat(float64(transaction.BillAmount), 'f', 6, 64),
			strconv.FormatFloat(float64(transaction.PrincipalAmount), 'f', 6, 64), strconv.FormatFloat(float64(transaction.InterestAmount), 'f', 6, 64),
			strconv.FormatFloat(float64(transaction.TotalFeeAmount), 'f', 6, 64), transaction.Status, transaction.PaymentMethod, strconv.Itoa(transaction.AutoDebetCounter),
			strings.Trim(transaction.CreatedAt, "T00:00:00Z"), strings.Trim(transaction.UpdatedAt, "T00:00:00Z"), strconv.FormatBool(transaction.IsHold), strconv.FormatBool(transaction.IsFstlPending),
			strconv.FormatBool(transaction.IsHstlPending), strconv.FormatBool(transaction.IsLaaPositif), strconv.FormatFloat(float64(transaction.PaymentAmount), 'f', 6, 64),
			strings.Trim(transaction.BillingGenDate, "T00:00:00Z"), strconv.FormatBool(transaction.IsOdaPositif)}
		stringData = append(stringData, record)
	}
	//
	return stringData, nil
=======
=======
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
func (u UseCase) GetTransactionByStatus(status string) ([]Transaction, error) {
	return u.repo.GetTransactionByStatus(status)
}

func (u UseCase) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error) {
	return u.repo.GetTransactionByStatusAndDate(req, input)
}

func (u UseCase) GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error) {
	return u.repo.GetTransactionByDate(req, input)
>>>>>>> parent of 276e577 (feat: add unit test layer usecase & controller)
}

func (u UseCase) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	return u.repo.GetAllTransactionByDate(start, end)

}

func (u UseCase) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	return u.repo.GetAllTransactionByDate(start, end)

}
