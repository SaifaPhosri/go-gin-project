package routes

import (
	"go-gin-api/handlers"
	"go-gin-api/repository"
	"go-gin-api/service"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	repo := repository.NewStudentRepository()
	svc := service.NewStudentService(repo)
	h := handlers.NewStudentHandler(svc)
	// student
	r.POST("/student", h.CreateStudent)
	r.GET("/student", h.GetAllStudent)
	r.GET("/student/:id", h.GetStudentById)
	r.PATCH("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)
}
