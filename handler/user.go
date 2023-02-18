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

// function login user
func (h *userHandler) Login(c *gin.Context) {
	// capture input from the user
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Login Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// map input from user to struct Login
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"error": err.Error()}
		// response error output
		respons := helper.APIResponse("Login Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// call function format user and save data to formatter
	formatter := user.FormatUser(loggedinUser, "tokennnn")

	// response success and send formatter data
	response := helper.APIResponse("Login Success!", http.StatusOK, "success", formatter)

	// sending json data
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// capture input from the user
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// call function FormatValidationError from helper
		errors := helper.FormatValidationError(err)
		// formatter error output
		errorMessage := gin.H{"error": errors}
		// response error output
		respons := helper.APIResponse("Email checking Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// map input from user to struct Check Email
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"error": "Server Error"}
		// response error output
		respons := helper.APIResponse("Email checking Failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// data format response
	data := gin.H {
		"is_available": isEmailAvailable,
	}

	// meta message for response
	metaMessage := "Email is not available"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	// response success and send formatter data
	respons := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	
	// sending json data
	c.JSON(http.StatusOK, respons)
}