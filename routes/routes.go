package routes

import (
	"github.com/anggakrnwn/user-auth-api/controllers"
	"github.com/anggakrnwn/user-auth-api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/api/register", controllers.Register)

	router.POST("/api/login", controllers.Login)

	router.GET("/api/users", middlewares.AuthMiddlewares(), controllers.FindUsers)

	router.POST("/api/users", middlewares.AuthMiddlewares(), controllers.CreteUser)

	router.GET("/api/users/:id", middlewares.AuthMiddlewares(), controllers.FindUserById)

	router.PATCH("/api/users/:id", middlewares.AuthMiddlewares(), controllers.UpdateUser)

	router.DELETE("/api/users/:id", middlewares.AuthMiddlewares(), controllers.DeleteUser)

	return router
}
