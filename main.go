package main

import (
	"fmt"
	"log"
	"net/http"
	"website-fundright/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// koneksi database
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_website_fundright?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to database is successful")

	// ambil nilai dari database
	var users []user.User
	db.Find(&users)

	// cetak nilai nama dari database
	for _, user := range users {
		fmt.Println(user.Name)
	}

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	// koneksi database
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_website_fundright?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to database is successful")

	// ambil nilai dari database
	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}