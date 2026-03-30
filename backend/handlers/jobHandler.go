package handlers

import (
	"database/sql"
	"job-tracker-backend/db"
	"job-tracker-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetJobs fetches all jobs from the database
func GetJobs(c *gin.Context) {
	jobs := []models.Job{}
	err := db.DB.Select(&jobs, "SELECT * FROM jobs")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

// CreateJob creates a new job entry
func CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set AppliedDate to now if empty
	if !job.AppliedDate.Valid {
		job.AppliedDate = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	// Keep InterviewDate NULL if empty
	if !job.InterviewDate.Valid {
		job.InterviewDate = sql.NullTime{
			Valid: false,
		}
	}

	// Insert job into DB
	res, err := db.DB.Exec(
		"INSERT INTO jobs (company, role, status, applied_date, interview_date) VALUES (?, ?, ?, ?, ?)",
		job.Company, job.Role, job.Status, job.AppliedDate, job.InterviewDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	job.ID = int(id)
	c.JSON(http.StatusOK, job)
}

// UpdateJob updates an existing job by ID
func UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure NullTime is properly set
	if !job.AppliedDate.Valid {
		job.AppliedDate = sql.NullTime{Valid: false}
	}
	if !job.InterviewDate.Valid {
		job.InterviewDate = sql.NullTime{Valid: false}
	}

	_, err := db.DB.Exec(
		"UPDATE jobs SET company=?, role=?, status=?, applied_date=?, interview_date=? WHERE id=?",
		job.Company, job.Role, job.Status, job.AppliedDate, job.InterviewDate, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteJob deletes a job by ID
func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM jobs WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
