package account

type UseCaseInterface interface {
	GetDataUser() ([]Account, error)
	GetDataUserById(id int) (Account, error)
	EditDataUser(id string, req *Account) (Account, error)
	DeleteDataUser(id string) (Account, error)
	CreateAccount(req *Account) (Account, error)
	Login(req *Account) (string, Account, error)
	SendEmail(email string) (Account, error)
	CompareVerificationCode(verificationCode *VerificationCodeRequest) (Account, error)
	EditPassword(id int, req *Account) (Account, error)
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
func (u UseCase) GetDataUserById(id int) (Account, error) {
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

func (u UseCase) SendEmail(email string) (Account, error) {
	result, err := u.Repo.SendEmail(email)
	if err != nil {
		return Account{}, err
	}
	return result, nil

}
func (u UseCase) CompareVerificationCode(verificationCode *VerificationCodeRequest) (Account, error) {
	result, err := u.Repo.CompareVerificationCode(verificationCode)
	if err != nil {
		return Account{}, err
	}
	return result, nil

}
func (u UseCase) EditPassword(id int, req *Account) (Account, error) {
	result, err := u.Repo.EditPassword(id, req)
	if err != nil {
		return Account{}, err
	}
	return result, nil
}
