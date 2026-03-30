package main

import (
	"fmt"
	"job-notifier/db"
	"job-notifier/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.Connect()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Start notification ticker
	go func() {
		ticker := time.NewTicker(1 * time.Minute) // demo every minute
		for range ticker.C {
			checkNotifications()
		}
	}()

	r.Run(":8081") // run notification service on different port
}

func checkNotifications() {
	jobs := []models.Job{}
	err := db.DB.Select(&jobs, "SELECT * FROM jobs")
	if err != nil {
		fmt.Println("Error fetching jobs:", err)
		return
	}

	now := time.Now()

	for _, job := range jobs {
		// Follow-up after applying
		if job.Status == "Applied" && !job.FollowUpSent && job.AppliedDate.Valid {
			if now.Sub(job.AppliedDate.Time).Hours() >= 24*7 {
				fmt.Println("Reminder: Follow up with", job.Company)
				db.DB.Exec("UPDATE jobs SET follow_up_sent=1 WHERE id=?", job.ID)
			}
		}

		// Follow-up after interview
		if job.Status == "Interview" && !job.InterviewFollowUpSent && job.InterviewDate.Valid {
			if now.Sub(job.InterviewDate.Time).Hours() >= 24*7 {
				fmt.Println("Reminder: Follow up on interview results at", job.Company)
				db.DB.Exec("UPDATE jobs SET interview_followup_sent=1 WHERE id=?", job.ID)
			}
		}
	}
}
