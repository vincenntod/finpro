package account

import (
	"golang/auth"
	"golang/module/account/dto"
	"golang/module/account/entities"
	"golang/packages"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseInterface interface {
	GetDataUser() ([]entities.Account, error)
	GetDataUserById(id string) (entities.Account, error)
	EditDataUser(id string, req *dto.EditDataUserRequest) (entities.Account, error)
	DeleteDataUser(id string) (entities.Account, error)
	CreateAccount(req *dto.CreateRequest) (entities.Account, error)
	Login(req *dto.LoginResponseRequest) (string, entities.Account, error)
	SendEmailForgotPassword(email string) (entities.Account, error)
	SendEmailRegistration(email string) (entities.Account, error)
	CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error)
	EditPassword(code string, req *dto.EditDataUserRequest) (entities.Account, error)
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
func (u UseCase) EditDataUser(id string, req *dto.EditDataUserRequest) (entities.Account, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)
	result, err := u.Repo.EditDataUser(id, req)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
func (u UseCase) CreateAccount(req *dto.CreateRequest) (entities.Account, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

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
func (u UseCase) Login(req *dto.LoginResponseRequest) (string, entities.Account, error) {
	result, err := u.Repo.Login(req)
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		return "", result, err
	}
	if err != nil {
		return "", result, err
	}
	expiredTime := time.Now().Add(time.Minute * 90)

	claims := &auth.JWT{
		Id:   result.Id,
		Name: result.Name,
		Role: result.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "golang",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenGenerate, err := token.SignedString(auth.JWT_KEY)
	return tokenGenerate, result, nil
}

func (u UseCase) SendEmailForgotPassword(email string) (entities.Account, error) {
	result, err := u.Repo.SendEmail(email)
	if err != nil {
		return entities.Account{}, err
	}
	if result.Id != 0 {
		packages.ConfigurateEmailRegistration(email)
	}
	return result, err
}
func (u UseCase) SendEmailRegistration(email string) (entities.Account, error) {
	result, err := u.Repo.SendEmail(email)
	if err != nil {
		return entities.Account{}, err
	}
	if result.Id == 0 {
		packages.ConfigurateEmailRegistration(email)
	}
	return result, err
}
func (u UseCase) CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error) {
	result, err := u.Repo.CompareVerificationCode(verificationCode)

	if err != nil {
		return entities.Account{}, err
	}
	return result, nil

}
func (u UseCase) EditPassword(id string, req *dto.EditDataUserRequest) (entities.Account, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)
	result, err := u.Repo.EditPassword(id, req)
	if err != nil {
		return entities.Account{}, err
	}
	return result, nil
}
