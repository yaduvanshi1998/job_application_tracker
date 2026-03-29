# job_application_tracker
While applying to many jobs, I needed a way to track applications, statuses, and follow-ups in one place, so I built a Job Application Tracker using React and Golang


1. Initialize Go Module
cd backend
go mod init job-tracker-backend
go get github.com/gin-gonic/gin
go get github.com/jmoiron/sqlx       # optional for DB
go get github.com/mattn/go-sqlite3  # if using SQLite