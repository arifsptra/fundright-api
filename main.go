package main

import (
	"log"
	"website-fundright/auth"
	"website-fundright/handler"
	"website-fundright/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect to database
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_website_fundright?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// error handling
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	
	router := gin.Default()

	// api version
	api := router.Group("/api/v1")

	// api endpoint for register
	api.POST("/users", userHandler.RegisterUser)
	
	// api endpoint for login
	api.POST("/sessions", userHandler.Login)

	// api endpoint for email checker
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	// api endpoint for upload avatar
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}