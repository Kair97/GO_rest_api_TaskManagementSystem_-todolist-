package routes

import (
	"try/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/tasks/:id", handlers.GetTaskById)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.PATCH("/tasks/:id", handlers.PatchTask)
	r.GET("/tasks", handlers.GetTasks)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

}
