package handler

import (
	"fmt"
	"net/http"
	"website-fundright/auth"
	"website-fundright/helper"
	"website-fundright/user"

	"github.com/gin-gonic/gin"
)

// this handler has the goal of mapping input from the user to the struct
type userHandler struct {
	userService user.Service
	authService auth.Service
}

// pass this struct as a service parameter
func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	// call jwt service
	token, err := h.authService.GenerateToken(newUser.ID)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Register Account is Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// call function format user and save data to formatter
	formatter := user.FormatUser(newUser, token)

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

	// call jwt service
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	// error handling
	if err != nil {
		// response error output
		respons := helper.APIResponse("Login Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// call function format user and save data to formatter
	formatter := user.FormatUser(loggedinUser, token)

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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// 
	file, err := c.FormFile("avatar")
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"is_uploaded": false}
		// response error output
		respons := helper.APIResponse("Avatar Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// for now user id is static value. later will use JWT
	userID := 1

	// path file name
	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	
	// save file
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"is_uploaded": false}
		// response error output
		respons := helper.APIResponse("Avatar Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// save avatar
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		// formatter error output
		errorMessage := gin.H{"is_uploaded": false}
		// response error output
		respons := helper.APIResponse("Avatar Failed to Upload!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, respons)
		return
	}
	
	// formatter error output
	errorMessage := gin.H{"is_uploaded": true}
	// response error output
	respons := helper.APIResponse("Avatar Success to Upload!", http.StatusOK, "success", errorMessage)
	c.JSON(http.StatusOK, respons)
}