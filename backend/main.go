package main

import (
	"job-tracker-backend/db"
	"job-tracker-backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.GET("/jobs", handlers.GetJobs)
	r.POST("/jobs", handlers.CreateJob)
	r.PUT("/jobs/:id", handlers.UpdateJob)
	r.DELETE("/jobs/:id", handlers.DeleteJob)

	r.Run(":8080")
}
