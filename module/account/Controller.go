package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"
)

type ControllerInterface interface {
	GetDataUser() (*dto.MessageResponse, error)
	GetDataUserById(id string) (*dto.MessageResponse, error)
	CreateAccount(req *dto.CreateRequest) (*dto.MessageResponse, error)
	EditDataUser(id string, req *dto.EditDataUserRequest) (*dto.MessageResponse, error)
	DeleteDataUser(id string) (*dto.MessageResponse, error)
	Login(req *dto.LoginResponseRequest) (string, *dto.LoginResponse, error)
	SendEmailForgotPassword(email string) (*dto.MessageResponse, error)
	SendEmailRegistration(email string) (*dto.MessageResponse, error)
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
	_, err := c.UseCase.CreateAccount(req)
	if err != nil {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Duplicate account email",
			Data:    nil,
		}
		return res, err
	}

	res := &dto.MessageResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Save Data",
		Data:    nil,
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

	data, err := c.UseCase.EditDataUser(id, req)
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

	token, data, err := c.UseCase.Login(req)
	if err != nil {
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
func (c Controller) SendEmailForgotPassword(email string) (*dto.MessageResponse, error) {
	data, err := c.UseCase.SendEmailForgotPassword(email)
	if data.Id == 0 {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Email not found",
			Data:    nil,
		}
		return res, nil
	}
	if err != nil {
		res := &dto.MessageResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed Send Email",
			Data:    nil,
		}
		return res, err
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
func (c Controller) SendEmailRegistration(email string) (*dto.MessageResponse, error) {
	data, err := c.UseCase.SendEmailRegistration(email)
	if data.Id != 0 {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Email Already Taken",
			Data:    nil,
		}
		return res, nil
	}
	if err != nil {
		res := &dto.MessageResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed Send Email",
			Data:    nil,
		}
		return res, err
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
	data, err := c.UseCase.CompareVerificationCode(verificationCode)

	CodeFromMap := entities.VerificationCodes[verificationCode.Email]
	CodeFromRequest := verificationCode.Code
	if CodeFromMap != CodeFromRequest {
		res := &dto.MessageResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid Verification Code",
			Data:    nil,
		}
		return res, err
	}
	if err != nil {
		res := &dto.MessageResponse{
			Code:    500,
			Status:  "Error",
			Message: "Internal Server Error",
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

	data, _ := c.UseCase.EditPassword(id, req)
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
