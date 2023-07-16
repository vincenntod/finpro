package account

import (
	"golang/auth"
	"golang/module/account/dto"
	"golang/module/account/entities"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetDataUser() ([]entities.Account, error)
	GetDataUserById(id string) (entities.Account, error)
	EditDataUser(id string, req *entities.Account) (entities.Account, error)
	DeleteDataUser(id string) (entities.Account, error)
	CreateAccount(req *entities.Account) (entities.Account, error)
	Login(req *entities.Account) (string, entities.Account, error)
	SendEmail(email string) (entities.Account, error)
	CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error)
	EditPassword(id string, req *entities.Account) (entities.Account, error)
}

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return Repository{db}
}

func (r Repository) GetDataUser() ([]entities.Account, error) {
	var account []entities.Account
	err := r.Db.Find(&account).Error
	return account, err
}

func (r Repository) GetDataUserById(id string) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Find(&account, id).Error
	return account, err

}
func (r Repository) EditDataUser(id string, req *entities.Account) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("UPDATE account SET name=?,phone=?,role=?,password=? WHERE id = ?", req.Name, req.Phone, req.Role, req.Password, id).Scan(&account).Raw("SELECT * FROM account WHERE id = ?", id).Scan(&account).Error
	return account, err
}
func (r Repository) DeleteDataUser(id string) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Where("id = ?", id).Delete(&account).Error
	return account, err
}

func (r Repository) CreateAccount(req *entities.Account) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("INSERT INTO account (name,phone,role,password,email) VALUES (?,?,?,?,?)", req.Name, req.Phone, req.Role, req.Password, req.Email).Scan(&account).Error
	return account, err
}

func (r Repository) Login(req *entities.Account) (string, entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("SELECT * FROM account WHERE email = ? ORDER BY account.id LIMIT 1", req.Email).Scan(&account).Error
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

func GenerateVerificationCode(email string) string {

	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(9000) + 1000)

	entities.VerificationCodes[email] = code
	return code
}
func (r Repository) SendEmail(email string) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("SELECT * FROM account WHERE email = ?", email).Scan(&account).Error
	return account, err

}

func (r Repository) CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("SELECT * FROM account WHERE email = ?", verificationCode.Email).Scan(&account).Error
	return account, err

}
func (r Repository) EditPassword(id string, req *entities.Account) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("UPDATE account SET password=? WHERE id = ?", req.Password, id).Scan(&account).Raw("SELECT * FROM account WHERE account.id = ?", id).Scan(&account).Error
	return account, err

}
