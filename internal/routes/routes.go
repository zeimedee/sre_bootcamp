package routes

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/students")
	{
		api.GET("/healthcheck")
		api.POST("/create")
		api.GET("/get/:id")
		api.GET("/get-all")
		api.PUT("/update/id")
		api.DELETE("/delete/:id")
	}

	return router
}
