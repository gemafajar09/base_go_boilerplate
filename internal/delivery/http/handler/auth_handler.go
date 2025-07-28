package handler

import (
	"net/http"

	"go-project/internal/delivery/http/response"
	"go-project/internal/delivery/http/validator"
	"go-project/internal/domain"
	"go-project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewAuthHandler(uc *usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{userUsecase: uc}
}

// Request body untuk register
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// POST /register
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to register user",
			Errors:  validator.PesanError(err),
		})
		return
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     domain.UserRole,
	}

	createdUser, err := h.userUsecase.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to register user",
			Errors:  validator.PesanError(err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    createdUser.Id,
		"name":  createdUser.Name,
		"email": createdUser.Email,
	})
}

// Request body untuk login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// POST /login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to login user",
			Errors:  validator.PesanError(err),
		})
		return
	}

	token, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Success: false,
			Message: "Invalid email or password",
			Errors:  validator.PesanError(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
