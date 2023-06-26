package account

import "golang.org/x/crypto/bcrypt"

type ControllerInterface interface {
	GetDataUser() (*ReadResponse, error)
	GetDataUserById(id string) (*ReadResponse, error)
	CreateAccount(req *CreateRequest) (*CreateResponse, error)
	EditDataUser(id string, req *EditDataUserRequest) (*CreateResponse, error)
	DeleteDataUser(id string) (*CreateResponse, error)
	Login(req *LoginResponseRequest) (string, *LoginResponse, error)
}

type Controller struct {
	UseCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		UseCase: useCase,
	}
}

type CreateResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Status  string                `json:"status"`
	Data    []AccountItemResponse `json:"data"`
}
type ReadResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Status  string                `json:"status"`
	Data    []AccountItemResponse `json:"data"`
}
type AccountItemResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
type LoginResponse struct {
	Message string                   `json:"message"`
	Code    int                      `json:"code"`
	Status  string                   `json:"status"`
	Data    []LoginResponseWithToken `json:"data"`
}
type LoginResponseWithToken struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (c Controller) GetDataUser() (*ReadResponse, error) {
	account, err := c.UseCase.GetDataUser()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{
		Status:  "OK",
		Message: "Berhasil Get Data",
		Code:    200,
	}

	for _, account := range account {
		item := AccountItemResponse{
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
func (c Controller) GetDataUserById(id string) (*ReadResponse, error) {
	account, err := c.UseCase.GetDataUserById(id)
	if err != nil {
		return nil, err
	}
	if account.Id == 0 {
		res := &ReadResponse{
			Code:    204,
			Status:  "OK",
			Message: "Data Dengan ID: " + id + " Kosong",
			Data:    nil,
		}
		return res, nil
	}
	res := &ReadResponse{
		Code:    200,
		Status:  "OK",
		Message: "Berhasil Get Data",
		Data: []AccountItemResponse{{
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
func (c Controller) CreateAccount(req *CreateRequest) (*CreateResponse, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := Account{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	_, err := c.UseCase.CreateAccount(&request)
	if err != nil {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Duplicate account name",
			Data:    nil,
		}
		return res, err
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Berhasil Simpan Data",
		Data: []AccountItemResponse{{
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
func (c Controller) EditDataUser(id string, req *EditDataUserRequest) (*CreateResponse, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := Account{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	_, err := c.UseCase.EditDataUser(id, &request)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Berhasil Edit Data",
		Data: []AccountItemResponse{{
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
func (c Controller) DeleteDataUser(id string) (*CreateResponse, error) {
	_, err := c.UseCase.DeleteDataUser(id)
	if err != nil {
		return nil, err
	}
	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Berhasil Delete Data",
	}
	return res, nil
}
func (c Controller) Login(req *LoginResponseRequest) (string, *LoginResponse, error) {
	request := Account{
		Email:    req.Email,
		Password: req.Password,
	}
	token, data, err := c.UseCase.Login(&request)
	if err != nil {
		res := &LoginResponse{
			Code:    401,
			Status:  "Error",
			Message: "Email atau Password Salah",
			Data:    nil,
		}
		return "", res, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password)); err != nil {
		res := &LoginResponse{
			Code:    401,
			Status:  "Error",
			Message: "Email atau Password Salah",
			Data:    nil,
		}
		return "", res, err
	}

	res := &LoginResponse{
		Code:    200,
		Status:  "OK",
		Message: "Login Berhasil",
		Data: []LoginResponseWithToken{{
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
