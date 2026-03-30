package models

import "time"

type Job struct {
	ID                    int       `db:"id" json:"id"`
	Company               string    `db:"company" json:"company"`
	Role                  string    `db:"role" json:"role"`
	Status                string    `db:"status" json:"status"`
	AppliedDate           time.Time `db:"applied_date" json:"applied_date"`
	InterviewDate         time.Time `db:"interview_date" json:"interview_date"`
	FollowUpSent          bool      `db:"follow_up_sent" json:"follow_up_sent"`
	InterviewFollowUpSent bool      `db:"interview_followup_sent" json:"interview_followup_sent"`
}
