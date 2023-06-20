package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	ctrl *Controller
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type CreateRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type LoginResponseRequest struct {
	Name     string `json:"name"`
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
	Email    string `json:"email"`
}

func (h RequestHandler) GetDataUser(c *gin.Context) {
	res, err := h.ctrl.GetDataUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h RequestHandler) GetDataUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.ctrl.GetDataUserById(id)
	if err != nil {
		c.JSON(500, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h RequestHandler) CreateAccount(c *gin.Context) {
	var account CreateRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.ctrl.CreateAccount(&account)
	if err != nil {
		c.JSON(400, gin.H{
			"Code":    400,
			"Status":  "Failed",
			"Message": "Duplicate Name",
		})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) EditDataUser(c *gin.Context) {
	var account EditDataUserRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.ctrl.EditDataUser(id, &account)
	if err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}
func (h RequestHandler) DeleteDataUser(c *gin.Context) {
	id := c.Param("id")
	res, err := h.ctrl.DeleteDataUser(id)
	if err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
	}
	c.JSON(200, res)
}
func (h RequestHandler) Login(c *gin.Context) {
	var account LoginResponseRequest

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	token, res, err := h.ctrl.Login(&account)
	if err != nil {
		c.JSON(401, gin.H{"code": 401,
			"message": "Username atau Password salah"})
		return
	}
	c.SetCookie("gin_cookie", token, 3600, "/", "localhost", false, true)
	c.JSON(200, res)
}
func (h RequestHandler) Logout(c *gin.Context) {
	c.SetCookie("gin_cookie", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "Berhasil Logout"})
}
