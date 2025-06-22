package controllers

import (
	"fmt"
	"net/http"

	"github.com/anggakrnwn/user-auth-api/database"
	"github.com/anggakrnwn/user-auth-api/helpers"
	"github.com/anggakrnwn/user-auth-api/models"
	"github.com/anggakrnwn/user-auth-api/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	fmt.Println("login berhasil")

	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	token := helpers.GenerateToken(user.Username)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.CreatedAt.String(),
			Token:     &token,
		},
	})
}
