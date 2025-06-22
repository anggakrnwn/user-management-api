package main

import (
	"fmt"

	"github.com/anggakrnwn/user-auth-api/config"
	"github.com/anggakrnwn/user-auth-api/database"
	"github.com/anggakrnwn/user-auth-api/routes"
)

func main() {
	fmt.Println("server dimulai")

	config.LoadEnv()

	database.InitDB()

	// router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world",
	// 	})
	// })

	// router.Run(":3001")

	r := routes.SetupRouter()
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
