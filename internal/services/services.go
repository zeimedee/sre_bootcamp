package services

import (
	"fmt"
	"log"

	"github.com/zeimedee/sre_bootcamp/internal/database"
	"github.com/zeimedee/sre_bootcamp/internal/models"
)

func CreateNewStudent(student models.Student) (models.Student, error) {
	database.DB.Db.Create(student)
	return student, nil
}

func GetStudent(id string) (models.Student, error) {
	students := models.Student{}
	result := database.DB.Db.Where("id = ?", id).Find(&students)
	if result.Error != nil {
		log.Println("error retrieving student:", result.Error)
		return students, result.Error
	}
	return students, nil
}

func GetAllStudent() ([]models.Student, error) {
	students := []models.Student{}
	result := database.DB.Db.Find(&students)
	if result.Error != nil {
		log.Println("error retrieving students:", result.Error)
		return students, result.Error
	}
	return students, nil
}

func UpdateStudent(student *models.Student, id string) error {
	result := database.DB.Db.Save(&student)
	if result.Error != nil {
		log.Println("error updating:", result.Error)
		return result.Error
	}
	return nil
}

func DeleteStudent(id string) error {
	var student models.Student
	res := database.DB.Db.Where("id = ?", id).Delete(&student)
	if res.Error != nil {
		log.Println("error deleting company: ", res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("no company found with the given id")
		return fmt.Errorf("no company found with id: %s", id)
	}
	return nil
}
