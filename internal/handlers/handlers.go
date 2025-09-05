package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zeimedee/sre_bootcamp/internal/models"
	"github.com/zeimedee/sre_bootcamp/internal/services"
)

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy new line"})
}

func CreateNewStudent(ctx *gin.Context) {
	student := new(models.Student)

	if err := ctx.ShouldBindBodyWithJSON(student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStudent := models.Student{
		Id:     uuid.NewString(),
		Name:   student.Name,
		Age:    student.Age,
		Course: student.Course,
	}
	_, err := services.CreateNewStudent(newStudent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": "failed to create student"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "true", "data": newStudent})
}

func GetStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := services.GetStudent(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "true", "data": data})
}

func GetAllStudent(ctx *gin.Context) {
	data, err := services.GetAllStudent()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "true", "data": data})

}

func UpdateNewStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	student := new(models.Student)
	if err := ctx.ShouldBindBodyWithJSON(student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.UpdateStudent(student, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "true", "msg": "student updated successfully"})
}

func DeleteStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	err := services.DeleteStudent(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "true", "msg": "student deleted successfully"})
}
