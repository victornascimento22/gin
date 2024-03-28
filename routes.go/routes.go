package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victornascimento22/gin.git/controllers"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run()
}
