package repository

import (
	"go-gin-api/config"
	"go-gin-api/models"

	"github.com/google/uuid"
)

type StudentRepository interface {
	CreateStudent(student models.Student) (models.Student, error)
	GetAllStudent() ([]models.Student, error)
	GetStudentById(id uuid.UUID) (models.Student, error)
	UpdateStudent(student models.Student) (models.Student, error)
	DeleteStudent(id uuid.UUID) error
}

type studentRepository struct{}

func NewStudentRepository() StudentRepository {
	return &studentRepository{}
}

func (r *studentRepository) CreateStudent(student models.Student) (models.Student, error) {
	if err := config.DB.Create(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentRepository) GetAllStudent() ([]models.Student, error) {
	var students []models.Student
	if err := config.DB.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) GetStudentById(id uuid.UUID) (models.Student, error) {
	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentRepository) UpdateStudent(student models.Student) (models.Student, error) {
	if err := config.DB.Model(&student).Where("id = ?", student.ID).Updates(student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentRepository) DeleteStudent(id uuid.UUID) error {
	if err := config.DB.Delete(&models.Student{}).Error; err != nil {
		return err
	}
	return nil
}
