package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"
)

type UseCaseInterface interface {
	GetDataUser() ([]entities.Account, error)
	GetDataUserById(id string) (entities.Account, error)
	EditDataUser(id string, req *entities.Account) (entities.Account, error)
	DeleteDataUser(id string) (entities.Account, error)
	CreateAccount(req *entities.Account) (entities.Account, error)
	Login(req *entities.Account) (string, entities.Account, error)
	SendEmail(email string) (entities.Account, error)
	CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error)
	EditPassword(code string, req *entities.Account) (entities.Account, error)
}

type UseCase struct {
	Repo RepositoryInterface
}

func NewUseCase(repo RepositoryInterface) UseCaseInterface {
	return &UseCase{
		Repo: repo,
	}
}

func (u UseCase) GetDataUser() ([]entities.Account, error) {
	result, err := u.Repo.GetDataUser()
	if err != nil {
		return []entities.Account{}, err
	}
	return result, nil
}
func (u UseCase) GetDataUserById(id string) (entities.Account, error) {
	result, err := u.Repo.GetDataUserById(id)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
func (u UseCase) EditDataUser(id string, req *entities.Account) (entities.Account, error) {
	result, err := u.Repo.EditDataUser(id, req)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
func (u UseCase) CreateAccount(req *entities.Account) (entities.Account, error) {
	result, err := u.Repo.CreateAccount(req)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
func (u UseCase) DeleteDataUser(id string) (entities.Account, error) {
	result, err := u.Repo.DeleteDataUser(id)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
func (u UseCase) Login(req *entities.Account) (string, entities.Account, error) {
	string, result, err := u.Repo.Login(req)
	if err != nil {
		return string, entities.Account{}, err
	}
	return string, result, nil
}

func (u UseCase) SendEmail(email string) (entities.Account, error) {
	result, err := u.Repo.SendEmail(email)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil

}
func (u UseCase) CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error) {
	result, err := u.Repo.CompareVerificationCode(verificationCode)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil

}
func (u UseCase) EditPassword(id string, req *entities.Account) (entities.Account, error) {
	result, err := u.Repo.EditPassword(id, req)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
