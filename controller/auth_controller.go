package controller

import (
	"gin_fleamarket/dto"
	"gin_fleamarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
}

type AuthController struct {
	services services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{services: service}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.services.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "faild to createuser"})
		return
	}
	ctx.Status(http.StatusCreated)
}
