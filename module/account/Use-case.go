package account

type UseCaseInterface interface {
	GetDataUser() ([]Account, error)
	GetDataUserById(id string) (Account, error)
	EditDataUser(id string, req *Account) (Account, error)
	DeleteDataUser(id string) (Account, error)
	CreateAccount(req *Account) (Account, error)
	Login(req *Account) (string, Account, error)
}

type UseCase struct {
	Repo UseCaseInterface
}

func NewUseCase(repo *Repository) *UseCase {
	if repo == nil {
		return nil
	}
	return &UseCase{
		Repo: repo,
	}
}

func (u UseCase) GetDataUser() ([]Account, error) {
	result, err := u.Repo.GetDataUser()
	if err != nil {
		return []Account{}, err
	}
	return result, nil
}
func (u UseCase) GetDataUserById(id string) (Account, error) {
	result, err := u.Repo.GetDataUserById(id)
	if err != nil {
		return Account{}, err
	}
	return result, nil
}
func (u UseCase) EditDataUser(id string, req *Account) (Account, error) {
	result, err := u.Repo.EditDataUser(id, req)
	if err != nil {
		return Account{}, err
	}
	return result, nil
}
func (u UseCase) CreateAccount(req *Account) (Account, error) {
	result, err := u.Repo.CreateAccount(req)
	if err != nil {
		return Account{}, err
	}
	return result, nil
}
func (u UseCase) DeleteDataUser(id string) (Account, error) {
	result, err := u.Repo.DeleteDataUser(id)
	if err != nil {
		return Account{}, err
	}
	return result, nil
}
func (u UseCase) Login(req *Account) (string, Account, error) {
	string, result, err := u.Repo.Login(req)
	if err != nil {
		return string, Account{}, err
	}
	return string, result, nil
}
