package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetDataUser() ([]entities.Account, error)
	GetDataUserById(id string) (entities.Account, error)
	EditDataUser(id string, req *dto.EditDataUserRequest) (entities.Account, error)
	DeleteDataUser(id string) (entities.Account, error)
	CreateAccount(req *dto.CreateRequest) (entities.Account, error)
	Login(req *dto.LoginResponseRequest) (entities.Account, error)
	SendEmail(email string) (entities.Account, error)
	CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error)
	EditPassword(id string, req *dto.EditDataUserRequest) (entities.Account, error)
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
func (r Repository) EditDataUser(id string, req *dto.EditDataUserRequest) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("UPDATE account SET name=?,phone=?,role=?,password=? WHERE id = ?", req.Name, req.Phone, req.Role, req.Password, id).Scan(&account).Raw("SELECT * FROM account WHERE id = ?", id).Scan(&account).Error
	return account, err
}
func (r Repository) DeleteDataUser(id string) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Where("id = ?", id).Delete(&account).Error
	return account, err
}

func (r Repository) CreateAccount(req *dto.CreateRequest) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("INSERT INTO account (name,phone,role,password,email) VALUES (?,?,?,?,?)", req.Name, req.Phone, req.Role, req.Password, req.Email).Scan(&account).Error
	return account, err
}

func (r Repository) Login(req *dto.LoginResponseRequest) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("SELECT * FROM account WHERE email = ? ORDER BY account.id LIMIT 1", req.Email).Scan(&account).Error
	if err != nil {
		return account, err
	}
	return account, nil
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
func (r Repository) EditPassword(id string, req *dto.EditDataUserRequest) (entities.Account, error) {
	var account entities.Account
	err := r.Db.Raw("UPDATE account SET password=? WHERE id = ?", req.Password, id).Scan(&account).Raw("SELECT * FROM account WHERE account.id = ?", id).Scan(&account).Error
	return account, err

}
