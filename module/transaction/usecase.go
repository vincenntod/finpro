package transaction

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) GetAllTransaction() ([]Transaction, error) {
	return u.repo.FindAll()

}

func (u UseCase) GetTransactionByStatus(status string) (Transaction, error) {
	return u.repo.FindByStatus(status)
}
