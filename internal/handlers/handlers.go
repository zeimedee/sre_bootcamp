package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	return
}

func CreateNewStudent(ctx *gin.Context) {}

func GetStudent(ctx *gin.Context) {}

func GetAllStudent(ctx *gin.Context) {}

func UpdateNewStudent(ctx *gin.Context) {}

func DeleteStudent(ctx *gin.Context) {}
