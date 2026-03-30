package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Job struct {
	ID                    int          `db:"id" json:"id"`
	Company               string       `db:"company" json:"company"`
	Role                  string       `db:"role" json:"role"`
	Status                string       `db:"status" json:"status"`
	AppliedDate           sql.NullTime `db:"applied_date" json:"applied_date"`
	InterviewDate         sql.NullTime `db:"interview_date" json:"interview_date"`
	FollowUpSent          bool         `db:"follow_up_sent" json:"follow_up_sent"`
	InterviewFollowUpSent bool         `db:"interview_followup_sent" json:"interview_followup_sent"`
}

func nullTime(s string) sql.NullTime {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil || s == "" {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: t, Valid: true}
}

func formatTime(t sql.NullTime) *string {
	if !t.Valid {
		return nil
	}
	s := t.Time.Format(time.RFC3339)
	return &s
}

func (j Job) MarshalJSON() ([]byte, error) {
	type Alias Job
	return json.Marshal(&struct {
		AppliedDate   *string `json:"applied_date"`
		InterviewDate *string `json:"interview_date"`
		*Alias
	}{
		AppliedDate:   formatTime(j.AppliedDate),
		InterviewDate: formatTime(j.InterviewDate),
		Alias:         (*Alias)(&j),
	})
}

func (j *Job) UnmarshalJSON(data []byte) error {
	type Alias struct {
		ID                    int    `json:"id"`
		Company               string `json:"company"`
		Role                  string `json:"role"`
		Status                string `json:"status"`
		AppliedDate           string `json:"applied_date"`
		InterviewDate         string `json:"interview_date"`
		FollowUpSent          bool   `json:"follow_up_sent"`
		InterviewFollowUpSent bool   `json:"interview_followup_sent"`
	}
	var a Alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	j.ID, j.Company, j.Role, j.Status = a.ID, a.Company, a.Role, a.Status
	j.FollowUpSent, j.InterviewFollowUpSent = a.FollowUpSent, a.InterviewFollowUpSent
	j.AppliedDate, j.InterviewDate = nullTime(a.AppliedDate), nullTime(a.InterviewDate)
	return nil
}
