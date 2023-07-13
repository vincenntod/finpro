package dto

// Request
type CreateRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type LoginResponseRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
type EditDataUserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
type VerificationCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// Controller Response
type MessageResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Status  string                `json:"status"`
	Data    []AccountItemResponse `json:"data"`
}
type LoginResponse struct {
	Message string                   `json:"message"`
	Code    int                      `json:"code"`
	Status  string                   `json:"status"`
	Data    []LoginResponseWithToken `json:"data"`
}

type AccountItemResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
type LoginResponseWithToken struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Token string `json:"token"`
}
