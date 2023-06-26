package account

import (
	"errors"
	"fmt"
	"golang/auth"
	"golang/model"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
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

var VerificationCodes = make(map[string]string)

func GenerateVerificationCode(email string) string {

	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(9000) + 1000)

	VerificationCodes[email] = code
	return code
}
func SendEmail(c *gin.Context) {
	var account Account
	email := c.Param("email")
	model.DB.Raw("SELECT * FROM account WHERE email = ?", email).Scan(&account)
	if account.Id == 0 {
		c.JSON(500, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Email not found",
			"data":    nil,
		})
		return
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "briceria123@gmail.com"
	password := "fjuyqpqqnrkzaatr"
	rand := GenerateVerificationCode(email)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	subject := "Subject: [Verification Code] Forgot Password Account App Dashboard DDB Ceria\n"

	mainMessage1 := "<body marginheight='0' topmargin='0' marginwidth='0' style='margin: 0px; background-color: #f2f3f8;' leftmargin='0'><table cellspacing='0' border='0' cellpadding='0' width='100%' bgcolor='#f2f3f8'style='@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;'><tr><td><table style='background-color: #f2f3f8; max-width:670px;  margin:0 auto;' width='100%' border='0'align='center' cellpadding='0' cellspacing='0'><tr><td style='height:80px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><a href='https://bri.co.id/web/ceria' title='logo' target='_blank'><img width='60' src='https://play-lh.googleusercontent.com/tpsB_EJ4_p3Ljh7LwhNWg6ysAH8GoDzDIcZwIWTP9SX1HsVjPflGP_iUK4IWGZOulDk=w480-h960-rw' title='logo' alt='logo'></a></td></tr><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td><table width='95%' border='0' align='center' cellpadding='0' cellspacing='0'style='max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);'><tr><td style='height:40px;'>&nbsp;</td></tr><tr><td style='padding:0 35px;'><h1 style='color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;'>You have requested to reset your password</h1><spanstyle='display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;'></span><p style='color:#455056; font-size:15px;line-height:24px; margin:0;'>We cannot simply send you your old password. Please copy your verification code. To reset your password</p><a href='javascript:void(0);'style='background:#20e277;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;'>"
	mainMessage2 := "</a></td></tr><tr><td style='height:40px;'>&nbsp;</td></tr></table></td><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><p style='font-size:14px; color:rgba(69, 80, 86, 0.7411764705882353); line-height:18px; margin:0 0 0;'>&copy; <strong>https://bri.co.id/web/ceria</strong></p></td></tr><tr><td style='height:80px;'>&nbsp;</td></tr></table></td></tr></table></body>"

	body := mainMessage1 + rand + mainMessage2
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"Code":    200,
		"Status":  "OK",
		"Message": "Success Send Email",
		"data": gin.H{
			"id":    &account.Id,
			"email": &account.Email,
		},
	})
}

type VerificationCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func CompareVerificationCode(c *gin.Context) {
	var verificationCodeRequest VerificationCodeRequest
	var account Account

	if err := c.ShouldBindJSON(&verificationCodeRequest); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
	}
	CodeFromMap := VerificationCodes[verificationCodeRequest.Email]
	CodeFromRequest := verificationCodeRequest.Code
	if CodeFromMap != CodeFromRequest {
		c.JSON(401, gin.H{
			"code":    401,
			"status":  "Error",
			"message": "Invalid Code",
		})
		return
	}
	model.DB.Raw("Select * FROM account WHERE email = ?", verificationCodeRequest.Email).Scan(&account)
	if account.Id == 0 {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Email tidak terdaftar",
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Success Verification Code",
		"data": gin.H{
			"id":    &account.Id,
			"name":  &account.Name,
			"email": &account.Email,
			"role":  &account.Role,
			"phone": &account.Phone,
		},
	})
}
func EditPassword(c *gin.Context) {
	var account Account
	queryUrl := c.Request.URL.Query()
	id := queryUrl.Get("id")
	if err := model.DB.Find(&account, id).Error; err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request Id Not Found",
		})
		return
	}
	Code := c.Request.Header.Get("VerificationCode")
	if Code == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request Verification Code nil",
		})
		return
	}

	CodeFromMap := VerificationCodes[account.Email]
	CodeFromRequest := Code
	if CodeFromMap != CodeFromRequest {
		c.JSON(401, gin.H{
			"code":    401,
			"status":  "Error",
			"message": "Invalid Code",
		})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
	}
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((account.Password)), bcrypt.DefaultCost)
	account.Password = string(HashPassword)
	result := model.DB.Where("id = ?", id).Updates(&account).Find(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Failed to update account",
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Success Update Password",
		"data": gin.H{
			"id":    &account.Id,
			"email": &account.Email,
		},
	})
}
