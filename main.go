package main

import (
	"log"
	"net/http"
	"strings"
	"website-fundright/auth"
	"website-fundright/campaign"
	"website-fundright/handler"
	"website-fundright/helper"
	"website-fundright/payment"
	"website-fundright/transaction"
	"website-fundright/user"

	"github.com/dgrijalva/jwt-go"
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

	campaignRepository := campaign.NewRepository(db)

	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)

	campaignService := campaign.NewService(campaignRepository)

	paymentService := payment.NewService()

	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	campaignHandler := handler.NewCampaignHandler(campaignService)

	transactionHandler := handler.NewTransactionHandler(transactionService)
	
	router := gin.Default()

	// route for image directories
	router.Static("/images/avatars", "./images/avatars")
	router.Static("/images/campaigns", "./images/campaigns")

	// api version
	api := router.Group("/api/v1")

	// api endpoint for register
	api.POST("/users", userHandler.RegisterUser)
	
	// api endpoint for login
	api.POST("/sessions", userHandler.Login)

	// api endpoint for email checker
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	// api endpoint for upload avatar
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// api endpoint for get campaigns
	api.GET("/campaigns", campaignHandler.GetCampaigns)

	// api endpoint for get campaign by campaign id
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)

	// api endpoint for create new campaign
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)

	// api endpoint for update campaign
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	// api endpoint for post campaign image
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	// api endpoint for get transaction campaign
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)

	// api endpoint for get user transaction
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransaction)

	// api endpoint for create transaction
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	// api endpoint for create transaction notification
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run()
}

// function authentication middleware
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		// get context header "Authorization"
		authHeader := c.GetHeader("Authorization")
		// if first word in authHeader is not "Bearer"
		if !strings.Contains(authHeader, "Bearer") {
			// response error output
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// abort program with status json
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	
		// get token
		// declare token string
		tokenString := ""
		// default token format "Bearer thisistoken"
		tokenArray := strings.Split(authHeader, " ")
		// for this i get only token
		if len(tokenArray) == 2 {
			// initialize token string
			tokenString = tokenArray[1]
		}

		// validate token 
		token, err := authService.ValidateToken(tokenString)
		// error handling
		if err != nil {
			// response error output
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// abort program with status json
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload with jwt
		payload, ok := token.Claims.(jwt.MapClaims)
		// error handling
		if !ok || !token.Valid {
			// response error output
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// abort program with status json
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get user id
		userID := int(payload["user_id"].(float64))

		// get user by id
		user, err := userService.GetUserByID(userID)
		if err != nil {
			// response error output
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// abort program with status json
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// set context value is user
		c.Set("currentUser", user)
	}
}