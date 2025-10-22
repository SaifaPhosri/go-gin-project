package routes

import (
	"go-gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// student
	r.POST("/student", handlers.CreateStudent)
	r.GET("/student", handlers.GetAllStudent)
	r.GET("/student/:id", handlers.GetStudentById)
	r.PATCH("/student/:id", handlers.UpdateStudent)
	r.DELETE("/student/:id", handlers.DeleteStudent)
}