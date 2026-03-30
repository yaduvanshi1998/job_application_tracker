package handlers

import (
	"job-tracker-backend/db"
	"job-tracker-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {
	jobs := []models.Job{}
	db.DB.Select(&jobs, "SELECT * FROM jobs")
	c.JSON(http.StatusOK, jobs)
}

func CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if job.AppliedDate.IsZero() {
		job.AppliedDate = time.Now()
	}

	res, _ := db.DB.Exec("INSERT INTO jobs (company, role, status, applied_date, interview_date) VALUES (?, ?, ?, ?, ?)",
		job.Company, job.Role, job.Status, job.AppliedDate, job.InterviewDate)
	id, _ := res.LastInsertId()
	job.ID = int(id)
	c.JSON(http.StatusOK, job)
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Exec("UPDATE jobs SET company=?, role=?, status=?, applied_date=?, interview_date=? WHERE id=?",
		job.Company, job.Role, job.Status, job.AppliedDate, job.InterviewDate, id)

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	db.DB.Exec("DELETE FROM jobs WHERE id=?", id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
