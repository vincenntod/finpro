package account

import (
	"golang/auth"
	"golang/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetDataUser() ([]Account, error)
	GetDataUserById(id string) (Account, error)
	EditDataUser(id string, req *Account) (Account, error)
	DeleteDataUser(id string) (Account, error)
	CreateAccount(req *Account) (Account, error)
	Login(req *Account) (string, Account, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetDataUser() ([]Account, error) {
	var account []Account
	err := model.DB.Find(&account).Error
	return account, err
}

func (r Repository) GetDataUserById(id string) (Account, error) {
	var account Account
	err := model.DB.Find(&account, id).Error
	return account, err

}
func (r Repository) EditDataUser(id string, req *Account) (Account, error) {
	var account Account
	err := model.DB.Where("id = ?", id).Updates(&req).Error
	return account, err
}
func (r Repository) DeleteDataUser(id string) (Account, error) {
	var account Account
	err := model.DB.Where("id = ?", id).Delete(&account).Error
	return account, err
}

func (r Repository) CreateAccount(req *Account) (Account, error) {
	var account Account
	err := model.DB.Create(&req).Error
	return account, err
}

func (r Repository) Login(req *Account) (string, Account, error) {
	var account Account

	// var getAcount Account
	err := model.DB.Where("email = ?", req.Email).First(&account).Error
	if err != nil {
		return "", account, err
	}
	expiredTime := time.Now().Add(time.Minute * 90)

	claims := &auth.JWT{
		Id:   account.Id,
		Name: account.Name,
		Role: account.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "golang",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenGenerate, err := token.SignedString(auth.JWT_KEY)

	return tokenGenerate, account, nil
}
