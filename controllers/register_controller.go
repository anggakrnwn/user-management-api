package controllers

import (
	"net/http"

	"github.com/anggakrnwn/user-auth-api/database"
	"github.com/anggakrnwn/user-auth-api/helpers"
	"github.com/anggakrnwn/user-auth-api/models"
	"github.com/anggakrnwn/user-auth-api/structs"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "validation erros",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	hashed, err := helpers.HashedPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: hashed,
	}

	if err := database.DB.Create(&user).Error; err != nil {

		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "failed to create user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
