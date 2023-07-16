package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"
	"net/smtp"

	"golang.org/x/crypto/bcrypt"
)

type ControllerInterface interface {
	GetDataUser() (*dto.MessageResponse, error)
	GetDataUserById(id string) (*dto.MessageResponse, error)
	CreateAccount(req *dto.CreateRequest) (*dto.MessageResponse, error)
	EditDataUser(id string, req *dto.EditDataUserRequest) (*dto.MessageResponse, error)
	DeleteDataUser(id string) (*dto.MessageResponse, error)
	Login(req *dto.LoginResponseRequest) (string, *dto.LoginResponse, error)
	SendEmail(email string) (*dto.MessageResponse, error)
	SendEmailRegister(email string) (*dto.MessageResponse, error)
	CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (*dto.MessageResponse, error)
	EditPassword(id string, code string, req *dto.EditDataUserRequest) (*dto.MessageResponse, error)
}

type Controller struct {
	UseCase UseCaseInterface
}

func NewController(useCase UseCaseInterface) ControllerInterface {
	return Controller{
		UseCase: useCase,
	}
}

func (c Controller) GetDataUser() (*dto.MessageResponse, error) {
	account, err := c.UseCase.GetDataUser()
	if err != nil {
		res := &dto.MessageResponse{
			Status:  "Error",
			Message: "Internal Server Error",
			Code:    400,
			Data:    nil,
		}
		return res, nil
	}

	res := &dto.MessageResponse{
		Status:  "OK",
		Message: "Success Get Data",
		Code:    200,
	}

	for _, account := range account {
		item := dto.AccountItemResponse{
			Id:    account.Id,
			Name:  account.Name,
			Role:  account.Role,
			Email: account.Email,
			Phone: account.Phone,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) GetDataUserById(id string) (*dto.MessageResponse, error) {
	account, err := c.UseCase.GetDataUserById(id)
	if err != nil {
		return nil, err
	}
	if account.Id == 0 {
		res := &dto.MessageResponse{
			Code:    204,
			Status:  "OK",
			Message: "Data With ID: " + id + " Empty",
			Data:    nil,
		}
		return res, nil
	}
	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Get Data",
		Data: []dto.AccountItemResponse{{
			Id:    account.Id,
			Name:  account.Name,
			Role:  account.Role,
			Email: account.Email,
			Phone: account.Phone,
		},
		},
	}
	return res, nil
}
func (c Controller) CreateAccount(req *dto.CreateRequest) (*dto.MessageResponse, error) {
	if req.Name == "" || req.Phone == "" || req.Role == "" || req.Password == "" || req.Email == "" {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Field not complete",
			Data:    nil,
		}
		return res, nil
	}
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := entities.Account{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	_, err := c.UseCase.CreateAccount(&request)
	if err != nil {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Duplicate account name",
			Data:    nil,
		}
		return res, err
	}

	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Save Data",
		Data: []dto.AccountItemResponse{{
			Id:    request.Id,
			Name:  request.Name,
			Role:  request.Role,
			Email: request.Email,
			Phone: request.Phone,
		},
		},
	}
	return res, nil

}
func (c Controller) EditDataUser(id string, req *dto.EditDataUserRequest) (*dto.MessageResponse, error) {
	if req.Name == "" || req.Phone == "" || req.Role == "" || req.Password == "" {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Field not complete",
			Data:    nil,
		}
		return res, nil
	}
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := entities.Account{
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	data, err := c.UseCase.EditDataUser(id, &request)
	if err != nil {
		return nil, err
	}

	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Edit Data",
		Data: []dto.AccountItemResponse{{
			Id:    data.Id,
			Name:  data.Name,
			Role:  data.Role,
			Email: data.Email,
			Phone: data.Phone,
		},
		},
	}
	return res, nil
}
func (c Controller) DeleteDataUser(id string) (*dto.MessageResponse, error) {
	_, err := c.UseCase.DeleteDataUser(id)
	if err != nil {
		return nil, err
	}
	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Delete Data",
	}
	return res, nil
}
func (c Controller) Login(req *dto.LoginResponseRequest) (string, *dto.LoginResponse, error) {
	request := entities.Account{
		Email:    req.Email,
		Password: req.Password,
	}
	token, data, err := c.UseCase.Login(&request)
	if err != nil {
		res := &dto.LoginResponse{
			Code:    401,
			Status:  "Error",
			Message: "Email atau Password Salah",
			Data:    nil,
		}
		return "", res, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password)); err != nil {
		res := &dto.LoginResponse{
			Code:    401,
			Status:  "Error",
			Message: "Email atau Password Salah",
			Data:    nil,
		}
		return "", res, err
	}

	res := &dto.LoginResponse{
		Code:    200,
		Status:  "OK",
		Message: "Login Success",
		Data: []dto.LoginResponseWithToken{{
			Id:    data.Id,
			Name:  data.Name,
			Email: data.Email,
			Role:  data.Role,
			Token: token,
		},
		},
	}
	return token, res, nil
}
func (c Controller) SendEmail(email string) (*dto.MessageResponse, error) {
	data, _ := c.UseCase.SendEmail(email)
	if data.Id == 0 {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Email not found",
			Data:    nil,
		}
		return res, nil
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

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		res := &dto.MessageResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed Send Email",
			Data:    nil,
		}
		return res, nil
	}

	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Send Email",
		Data: []dto.AccountItemResponse{{
			Id:    data.Id,
			Email: data.Email,
			Name:  data.Name,
			Role:  data.Role,
			Phone: data.Phone,
		},
		},
	}
	return res, nil
}
func (c Controller) SendEmailRegister(email string) (*dto.MessageResponse, error) {
	data, _ := c.UseCase.SendEmail(email)
	if data.Id != 0 {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Email Already Taken",
			Data:    nil,
		}
		return res, nil
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "briceria123@gmail.com"
	password := "fjuyqpqqnrkzaatr"
	rand := GenerateVerificationCode(email)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	subject := "Subject: [Verification Code] Registrasion Account App Dashboard DDB Ceria\n"

	mainMessage1 := "<body marginheight='0' topmargin='0' marginwidth='0' style='margin: 0px; background-color: #f2f3f8;' leftmargin='0'><table cellspacing='0' border='0' cellpadding='0' width='100%' bgcolor='#f2f3f8'style='@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;'><tr><td><table style='background-color: #f2f3f8; max-width:670px;  margin:0 auto;' width='100%' border='0'align='center' cellpadding='0' cellspacing='0'><tr><td style='height:80px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><a href='https://bri.co.id/web/ceria' title='logo' target='_blank'><img width='60' src='https://play-lh.googleusercontent.com/tpsB_EJ4_p3Ljh7LwhNWg6ysAH8GoDzDIcZwIWTP9SX1HsVjPflGP_iUK4IWGZOulDk=w480-h960-rw' title='logo' alt='logo'></a></td></tr><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td><table width='95%' border='0' align='center' cellpadding='0' cellspacing='0'style='max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);'><tr><td style='height:40px;'>&nbsp;</td></tr><tr><td style='padding:0 35px;'><h1 style='color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;'>You have requested to registration your account</h1><spanstyle='display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;'></span><p style='color:#455056; font-size:15px;line-height:24px; margin:0;'>Please copy your verification code. To registration your password</p><a href='javascript:void(0);'style='background:#20e277;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;'>"
	mainMessage2 := "</a></td></tr><tr><td style='height:40px;'>&nbsp;</td></tr></table></td><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><p style='font-size:14px; color:rgba(69, 80, 86, 0.7411764705882353); line-height:18px; margin:0 0 0;'>&copy; <strong>https://bri.co.id/web/ceria</strong></p></td></tr><tr><td style='height:80px;'>&nbsp;</td></tr></table></td></tr></table></body>"

	body := mainMessage1 + rand + mainMessage2
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		res := &dto.MessageResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed Send Email",
			Data:    nil,
		}
		return res, nil
	}

	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Send Email",
		Data:    nil,
	}
	return res, nil
}
func (c Controller) CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (*dto.MessageResponse, error) {
	data, _ := c.UseCase.CompareVerificationCode(verificationCode)

	CodeFromMap := entities.VerificationCodes[verificationCode.Email]
	CodeFromRequest := verificationCode.Code
	if CodeFromMap != CodeFromRequest {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid Verification Code",
			Data:    nil,
		}
		return res, nil
	}
	if data.Id == 0 {
		res := &dto.MessageResponse{
			Code:    200,
			Status:  "OK",
			Message: "Success Verification Code",
			Data:    nil,
		}
		return res, nil
	}
	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Verification Code",
		Data: []dto.AccountItemResponse{{
			Id:    data.Id,
			Email: data.Email,
			Name:  data.Name,
			Role:  data.Role,
			Phone: data.Phone,
		},
		},
	}
	return res, nil

}
func (c Controller) EditPassword(id string, code string, req *dto.EditDataUserRequest) (*dto.MessageResponse, error) {
	result, err := c.UseCase.GetDataUserById(id)
	if err != nil {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Account Not Found",
			Data:    nil,
		}
		return res, nil
	}
	CodeFromMap := entities.VerificationCodes[result.Email]
	CodeFromRequest := code

	if CodeFromMap != CodeFromRequest {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid Verification Code",
			Data:    nil,
		}
		return res, nil
	}
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := entities.Account{
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	data, _ := c.UseCase.EditPassword(id, &request)
	if data.Id == 0 {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Failed to update account",
			Data:    nil,
		}
		return res, nil
	}

	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Update Password",
		Data: []dto.AccountItemResponse{{
			Id:    data.Id,
			Email: data.Email,
			Name:  data.Name,
			Role:  data.Role,
			Phone: data.Phone,
		},
		},
	}
	return res, nil
}
