package transaction

type UseCase struct {
	repo RepositoryInterface
}

type UseCaseInterface interface {
	GetAllTransaction() ([]Transaction, error)
	GetTransactionByStatus(status string) ([]Transaction, error)
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

func (u UseCase) GetTransactionByStatus(status string) ([]Transaction, error) {
	return u.repo.GetTransactionByStatus(status)
}

func (u UseCase) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error) {
	return u.repo.GetTransactionByStatusAndDate(req, input)
}

func (u UseCase) GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error) {
	return u.repo.GetTransactionByDate(req, input)

}
