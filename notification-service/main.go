package main

import (
	"fmt"
	"job-notifier/db"
	"job-notifier/models"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

var (
	notifications []Notification
	notifMu       sync.Mutex
	notifCounter  int
)

func addNotification(msg string) {
	notifMu.Lock()
	defer notifMu.Unlock()
	notifCounter++
	notifications = append(notifications, Notification{ID: notifCounter, Message: msg})
}

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

	go func() {
		// checkNotifications() // runs immediately on start
		ticker := time.NewTicker(3 * time.Minute)
		for range ticker.C {
			checkNotifications()
		}
	}()

	r.GET("/notifications", func(c *gin.Context) {
		notifMu.Lock()
		defer notifMu.Unlock()
		c.JSON(200, notifications)
	})

	r.DELETE("/notifications/:id", func(c *gin.Context) {
		id := c.Param("id")
		notifMu.Lock()
		defer notifMu.Unlock()
		for i, n := range notifications {
			if fmt.Sprintf("%d", n.ID) == id {
				notifications = append(notifications[:i], notifications[i+1:]...)
				break
			}
		}
		c.JSON(200, gin.H{"ok": true})
	})

	r.Run(":8081")
}

func checkNotifications() {
	jobs := []models.Job{}
	if err := db.DB.Select(&jobs, "SELECT * FROM jobs"); err != nil {
		return
	}

	now := time.Now()
	for _, job := range jobs {
		if job.Status == "Applied" && !job.FollowUpSent && job.AppliedDate.Valid {
			if now.Sub(job.AppliedDate.Time).Hours() >= 24*7 {
				addNotification("Follow up with " + job.Company)
				db.DB.Exec("UPDATE jobs SET follow_up_sent=1 WHERE id=?", job.ID)
			}
		}
		if job.Status == "Interview" && !job.InterviewFollowUpSent && job.InterviewDate.Valid {
			if now.Sub(job.InterviewDate.Time).Hours() >= 24*7 {
				addNotification("Follow up on interview results at " + job.Company)
				db.DB.Exec("UPDATE jobs SET interview_followup_sent=1 WHERE id=?", job.ID)
			}
		}
	}
}
