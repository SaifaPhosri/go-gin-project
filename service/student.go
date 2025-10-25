package service

import (
	"errors"
	"go-gin-api/models"
	"go-gin-api/repository"

	"github.com/google/uuid"
)

type StudentService interface {
	CreateStudent(student models.Student) (models.Student, error)
	GetAllStudent() ([]models.Student, error)
	GetStudentById(id uuid.UUID) (models.Student, error)
	UpdateStudent(student models.Student) (models.Student, error)
	DeleteStudent(id uuid.UUID) error
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(student models.Student) (models.Student, error) {
	if student.FirstName == "" || student.LastName == "" {
		return models.Student{}, errors.New("firstName and lastName cannot be empty")
	}

	if student.StudentCode == "" {
		return models.Student{}, errors.New("studentCode cannot be empty")
	}

	return s.repo.CreateStudent(student)
}

func (s *studentService) GetAllStudent() ([]models.Student, error) {
	return s.repo.GetAllStudent()
}

func (s *studentService) GetStudentById(id uuid.UUID) (models.Student, error) {
	return s.repo.GetStudentById(id)
}

func (s *studentService) UpdateStudent(student models.Student) (models.Student, error) {
	return s.repo.UpdateStudent(student)
}

func (s *studentService) DeleteStudent(id uuid.UUID) error {
	return s.repo.DeleteStudent(id)
}
