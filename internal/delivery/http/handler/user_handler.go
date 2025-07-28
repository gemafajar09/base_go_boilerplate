package handler

import (
	"go-project/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserHandler *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserHandler: uc}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	log.Print(exists)
	log.Print(userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User id not found"})
		return
	}
	c.JSON(200, gin.H{"userID": userID})
}
