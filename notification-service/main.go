package main

import (
	"fmt"
	"job-notifier/db"
	"time"
)

func main() {
	db.Connect()

	ticker := time.NewTicker(1 * time.Minute) // for demo, every minute
	for range ticker.C {
		checkNotifications()
	}
}

func checkNotifications() {
	jobs := []struct {
		ID                    int
		Company               string
		Status                string
		AppliedDate           time.Time
		InterviewDate         time.Time
		FollowUpSent          bool
		InterviewFollowUpSent bool
	}{}
	db.DB.Select(&jobs, "SELECT * FROM jobs")

	now := time.Now()

	for _, job := range jobs {
		// Follow-up after applying
		if job.Status == "Applied" && !job.FollowUpSent {
			if now.Sub(job.AppliedDate).Hours() >= 24*7 {
				fmt.Println("Reminder: Follow up with", job.Company)
				db.DB.Exec("UPDATE jobs SET follow_up_sent=1 WHERE id=?", job.ID)
			}
		}

		// Follow-up after interview
		if job.Status == "Interview" && !job.InterviewFollowUpSent {
			if now.Sub(job.InterviewDate).Hours() >= 24*7 {
				fmt.Println("Reminder: Follow up on interview results at", job.Company)
				db.DB.Exec("UPDATE jobs SET interview_followup_sent=1 WHERE id=?", job.ID)
			}
		}
	}
}
