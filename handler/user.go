package handler

import (
	"net/http"
	"website-fundright/helper"
	"website-fundright/user"

	"github.com/gin-gonic/gin"
)

// this handler has the goal of mapping input from the user to the struct
type userHandler struct {
	userService user.Service
}

// pass this struct as a service parameter
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// function register user
func (h *userHandler) RegisterUser(c *gin.Context) {
	// capture input from the user
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	// error handling
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Register Account is Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// map input from user to struct Register User
	newUser, err := h.userService.RegisterUser(input)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Register Account is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// call function format user and save data to formatter
	formatter := user.FormatUser(newUser, "tokennnn")

	// response success and send formatter data
	response := helper.APIResponse("Register Account is Success!", http.StatusOK, "success", formatter)

	// sending json data
	c.JSON(http.StatusOK, response)
}