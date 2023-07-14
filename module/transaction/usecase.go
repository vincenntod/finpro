package transaction

import "golang/module/transaction/entities"

type UseCase struct {
	Repo RepositoryInterface
}

type UseCaseInterface interface {
	GetAllTransaction(page int, size int) ([]entities.Transaction, error)
	GetAllTransactionByStatus(status string, page int, size int) ([]entities.Transaction, error)
	GetAllTransactionByDate(start string, end string, page int, size int) ([]entities.Transaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) ([]entities.Transaction, error)

	GetAllTransactionNoLimit() ([]entities.Transaction, error)
	GetAllTransactionByStatusNoLimit(status string) ([]entities.Transaction, error)
	GetAllTransactionByDateNoLimit(start string, end string) ([]entities.Transaction, error)
	GetAllTransactionByStatusDateNoLimit(status string, start string, end string) ([]entities.Transaction, error)
}

func NewUseCase(repo RepositoryInterface) UseCaseInterface {
	return UseCase{
		Repo: repo,
	}
}

func (u UseCase) GetAllTransaction(page int, size int) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransaction(page, size)

}

func (u UseCase) GetAllTransactionByStatus(status string, page int, size int) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionByStatus(status, page, size)
}

func (u UseCase) GetAllTransactionByDate(start string, end string, page int, size int) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionByDate(start, end, page, size)
}

func (u UseCase) GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionByStatusDate(status, start, end, page, size)
}

func (u UseCase) GetAllTransactionNoLimit() ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionNoLimit()

}

func (u UseCase) GetAllTransactionByStatusNoLimit(status string) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionByStatusNoLimit(status)
}

func (u UseCase) GetAllTransactionByDateNoLimit(start string, end string) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionByDateNoLimit(start, end)
}

func (u UseCase) GetAllTransactionByStatusDateNoLimit(status string, start string, end string) ([]entities.Transaction, error) {
	return u.Repo.GetAllTransactionByStatusDateNoLimit(status, start, end)
}
