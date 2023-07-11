package transaction

type UseCase struct {
	Repo RepositoryInterface
}

type UseCaseInterface interface {
	GetAllTransaction() ([]Transaction, error)
	GetAllTransactionByStatus(status string) ([]Transaction, error)
	GetAllTransactionByDate(start string, end string) ([]Transaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error)
	GetAllLimit(input FilterLimit) ([]Transaction, error, int64)

}

func NewUseCase(repo RepositoryInterface) UseCaseInterface {
	return UseCase{
		Repo: repo,
	}
}

func (u UseCase) GetAllTransaction() ([]Transaction, error) {
	return u.Repo.GetAllTransaction()

}

func (u UseCase) GetAllTransactionByStatus(status string) ([]Transaction, error) {
	return u.Repo.GetAllTransactionByStatus(status)
}

func (u UseCase) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	return u.Repo.GetAllTransactionByDate(start, end)
}

func (u UseCase) GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error) {
	return u.Repo.GetAllTransactionByStatusDate(status, start, end)
}

func (u UseCase) GetAllLimit(input FilterLimit) ([]Transaction, error, int64) {
	return u.Repo.GetAllLimit(input)
}

