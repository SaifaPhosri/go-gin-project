package handlers

import (
	"go-gin-api/config"
	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "create student success",
		"data": student,
	})
}

func GetAllStudent(c *gin.Context) {
	var student []models.Student

	if err := config.DB.Find(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":"find all student success",
		"total": len(student),
		"data": student,

	})
}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := config.DB.Find(&student, "id = ?", id ).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "fine student success",
		"data": student,
	})
}


func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := config.DB.Find(&student, "id = ?", id).Error; err != nil {
		c.JSON(500,gin.H{"error": err.Error()})
		return
	}

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&student).Updates(input).Error; err != nil {
		c.JSON(500,gin.H{"error": err.Error()})
		return
	}

	c.JSON(200,gin.H{
		"message": "update student success",
		"data": student,
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student

	if err := config.DB.Find(&student, "id = ?", id).Error; err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200,gin.H{"message": "delete student success"})
}