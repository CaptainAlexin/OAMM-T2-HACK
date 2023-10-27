package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// Initialize the Gin router.
	router := gin.Default()

	// Open the database connection.
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/OAMMT2HA?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the database tables.
	db.AutoMigrate(&User{})

	// Define the routes.
	router.GET("/users", func(ctx *gin.Context) {
		// Get all users from the database.
		var users []User
		db.Find(&users)

		// Respond with the users data in JSON format.
		ctx.JSON(200, users)
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		// Get the user with the specified ID from the database.
		var user User
		if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
			ctx.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		// Respond with the user data in JSON format.
		ctx.JSON(200, user)
	})

	router.POST("/users", func(ctx *gin.Context) {
		// Bind the JSON request body to a new user object.
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Create the new user in the database.
		db.Create(&user)

		// Respond with the new user data in JSON format.
		ctx.JSON(200, user)
	})

	router.PUT("/users/:id", func(ctx *gin.Context) {
		// Get the user with the specified ID from the database.
		var user User
		if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
			ctx.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		// Bind the JSON request body to the existing user object.
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Save the updated user in the database.
		db.Save(&user)

		// Respond with the updated user data in JSON format.
		ctx.JSON(200, user)
	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		// Get the user with the specified ID from the database.
		var user User
		if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
			ctx.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		// Delete the user from the database.
		db.Delete(&user)

		// Respond with a success message.
		ctx.JSON(200, gin.H{"message": "User deleted successfully!"})
	})

	// Start the Gin server.
	router.Run(":8080")
}
