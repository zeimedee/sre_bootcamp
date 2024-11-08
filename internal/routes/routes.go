package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeimedee/sre_bootcamp/internal/handlers"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1/students")
	{
		api.GET("/healthcheck", handlers.Healthcheck)
		api.POST("/create", handlers.CreateNewStudent)
		api.GET("/get/:id", handlers.GetStudent)
		api.GET("/get-all", handlers.GetAllStudent)
		api.PUT("/update/:id", handlers.UpdateNewStudent)
		api.DELETE("/delete/:id", handlers.DeleteStudent)
	}

	return router
}
