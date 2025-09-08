package handler

import (
	"fmt"
	"go-project/internal/delivery/http/response"
	"go-project/internal/delivery/http/validator"
	"go-project/internal/domain"
	"go-project/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AboutHandler struct {
	aboutUsecase *usecase.AboutUsecase
}

func NewAboutHandler(uc *usecase.AboutUsecase) *AboutHandler {
	return &AboutHandler{aboutUsecase: uc}
}

func (h *AboutHandler) GetAbout(c *gin.Context) {
	abouts, err := h.aboutUsecase.GetAbout()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, abouts)
}

type AboutRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (h *AboutHandler) CreateAbout(c *gin.Context) {
	var req AboutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to create about",
			Errors:  validator.PesanError(err),
		})
		return
	}
	about := domain.About{
		Title:   req.Title,
		Content: req.Content,
	}
	createabout, err := h.aboutUsecase.CreateAbout(about)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed to create about",
			Errors:  validator.PesanError(err),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"about": createabout,
	})
}

func (h *AboutHandler) EditAbout(c *gin.Context) {
	var id = c.Params.ByName("id")

	var num, err = strconv.Atoi(id)

	if err == nil {
		fmt.Println("data tidak boleh kosong")
	}

	editdata, err := h.aboutUsecase.EditAbout(num)

	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Success: false,
			Message: "Failed to edit about",
			Errors:  validator.PesanError(err),
		})
	}

	c.JSON(http.StatusFound, gin.H{
		"about": editdata,
	})
}

func (h *AboutHandler) UpdateAbout(c *gin.Context) {
	var id = c.Params.ByName("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req AboutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to update about",
			Errors:  validator.PesanError(err),
		})
		return
	}

	about := domain.About{
		Id:      uint(num),
		Title:   req.Title,
		Content: req.Content,
	}

	updatedabout, err := h.aboutUsecase.UpdateAbout(about)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed to update about",
			Errors:  validator.PesanError(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"about": updatedabout,
	})
}
