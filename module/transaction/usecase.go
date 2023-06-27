package transaction

type UseCase struct {
	repo RepositoryInterface
}

type UseCaseInterface interface {
	GetAllTransaction() ([]Transaction, error)
	GetAllTransactionByStatus(status string) ([]Transaction, error)
	GetAllTransactionByDate(start string, end string) ([]Transaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error)

	GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error)
	GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error)
}

func NewUseCase(repo RepositoryInterface) UseCaseInterface {
	return UseCase{
		repo: repo,
	}
}

func (u UseCase) GetAllTransaction() ([]Transaction, error) {
	return u.repo.GetAllTransaction()

}

func (u UseCase) GetAllTransactionByStatus(status string) ([]Transaction, error) {
	return u.repo.GetAllTransactionByStatus(status)
}

func (u UseCase) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	return u.repo.GetAllTransactionByDate(start, end)
}

func (u UseCase) GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error) {
	return u.repo.GetAllTransactionByStatusDate(status, start, end)
}

func (u UseCase) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error) {
	return u.repo.GetTransactionByStatusAndDate(req, input)
}

func (u UseCase) GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error) {
	return u.repo.GetTransactionByDate(req, input)
}
