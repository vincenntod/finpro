package account

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) GetDataUser() ([]Account, error) {
	return u.repo.GetDataUser()
}
func (u UseCase) GetDataUserById(id string) (Account, error) {
	return u.repo.GetDataUserById(id)
}
func (u UseCase) EditDataUser(id string, req *Account) (Account, error) {
	return u.repo.EditDataUser(id, req)
}
func (u UseCase) CreateAccount(req *Account) (Account, error) {
	return u.repo.CreateAccount(req)
}
func (u UseCase) DeleteDataUser(id string) (Account, error) {
	return u.repo.DeleteDataUser(id)
}
func (u UseCase) Login(req *Account) (string, Account, error) {
	return u.repo.Login(req)
}
