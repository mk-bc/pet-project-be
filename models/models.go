package models

import (
	"database/sql/driver"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

// adding enum values

type ApplicationStatus string

const (
	approved ApplicationStatus = "approved"
	pending  ApplicationStatus = "pending"
	rejected ApplicationStatus = "rejected"
)

func (a *ApplicationStatus) Scan(value interface{}) error {
	*a = ApplicationStatus(value.([]byte))
	return nil
}

func (a ApplicationStatus) Value() (driver.Value, error) {
	return string(a), nil
}

// db models

type SensitiveData struct {
	Email    string `gorm:"primary_key"`
	Password string
	Role     string
}

type Admin struct {
	gorm.Model
	AdminName  string
	AdminEmail string `gorm:"unique; not null"` //foreign key to sensitive data
}

type User struct {
	gorm.Model
	UserName  string
	UserEmail string //foreign key to userSensitiveData
	// UserPassword string
	UserImage   string
	PhoneNumber string
	Description string `gorm:"size:500"`
	Skills      string
	// Jobs        []*Job //`gorm:"many2many:job_applicants"`
	JobStatuses []*JobStatus
}

type Company struct {
	gorm.Model
	CompanyName  string
	CompanyEmail string `gorm:"not null"` //foreign key to companySenstiveData
	CompanyImage string
	Description  string
	// Jobs            []*Job //posts several jobs (has-many relationship)
	// instead of all the jobs stored without any order storing categories of jobs in a company is more meaningful.
	Categories []*JobCategory //job categories in which the company is hiring.
}

type Job struct {
	gorm.Model
	JobTitle       string
	JobDescription string `gorm:"size:500"`
	Salary         string
	CompanyID      uint `gorm:"not null"` //foreign key
	CategoryID     uint `gorm:"not null"` //foreign key
	// Applicants     []*User //`gorm:"many2many:job_applicants"`
}

type JobStatus struct {
	gorm.Model
	JobID             uint              `gorm:"not null"` //foreign key
	UserID            uint              `gorm:"not null"` //foreign key
	ApplicationStatus ApplicationStatus `sql:"type:application_status"`
}

type JobCategory struct {
	gorm.Model
	CategoryName string `gorm:"unique; not null"`
	Jobs         []*Job //jobs from different companies of same category (has-many relationship)
}

type SavedJob struct {
	gorm.Model
	JobID  uint `gorm:"not null"` //foreign key
	UserID uint `gorm:"not null"` //foreign key
}

func AddForeignKeys(db *gorm.DB) {

	db.Model(&Admin{}).AddForeignKey("admin_email", "sensitive_data(email)", "CASCADE", "CASCADE")

	db.Model(&User{}).AddForeignKey("user_email", "sensitive_data(email)", "CASCADE", "CASCADE")

	db.Model(&Company{}).AddForeignKey("company_email", "sensitive_data(email)", "CASCADE", "CASCADE")

	db.Model(&Job{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")

	db.Model(&Job{}).AddForeignKey("category_id", "job_categories(id)", "CASCADE", "CASCADE")

	db.Model(&JobStatus{}).AddForeignKey("job_id", "jobs(id)", "CASCADE", "CASCADE")

	db.Model(&JobStatus{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&SavedJob{}).AddForeignKey("job_id", "jobs(id)", "CASCADE", "CASCADE")

	db.Model(&SavedJob{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
