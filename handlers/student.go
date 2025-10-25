package handlers

import (
	"go-gin-api/config"
	"go-gin-api/models"
	"go-gin-api/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentHandler struct {
	service service.StudentService
}

func NewStudentHandler(service service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdStudent, err := h.service.CreateStudent(student)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "create student success",
		"data":    createdStudent,
	})
}

func (h *StudentHandler) GetAllStudent(c *gin.Context) {
	students, err := h.service.GetAllStudent()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "find all student success",
		"total":   len(students),
		"data":    students,
	})
}

func (h *StudentHandler) GetStudentById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"err": "invalid UUID"})
	}

	student, err := h.service.GetStudentById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "fine student success",
		"data":    student,
	})
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"err": "invalid UUID"})
	}
	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input.ID = id

	student, err := h.service.UpdateStudent(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "update student success",
		"data":    student,
	})
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(400, gin.H{"err": "invalid UUID"})
	}

	student := h.service.DeleteStudent(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "delete student success"})
}
