package config

import (
	"log"
	"os"
	"time"

	"github.com/Valgard/godotenv"
	"github.com/jinzhu/gorm"
	"github.com/mk-bc/pet-project-be/models"
)

func DBSetup() *gorm.DB {
	err := godotenv.Load("/home/manjunath/Desktop/learning/pet-project/pet-project-be/.env")
	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}
	db_details := os.Getenv("db_details")
	db, err := gorm.Open("postgres", db_details)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// debug
	// db.LogMode(true)
	return db
}

func DBStart() *gorm.DB {
	db := DBSetup()

	// adding custom enum type to the database
	if err := db.Exec("CREATE TYPE application_status AS ENUM ('approved', 'pending', 'rejected');").Error; err != nil {
		log.Printf("Error creating application_status enum: %v", err)
	}
	// creating database
	db.DropTableIfExists(&models.JobStatus{})
	db.DropTableIfExists(&models.SavedJob{})
	db.DropTableIfExists(&models.Job{}) //with foreign keys

	db.DropTableIfExists(&models.Admin{}, &models.User{}, &models.UserSensitiveData{}, &models.Company{}, &models.CompanySensitiveData{}, &models.Job{}, &models.JobCategory{}, &models.JobStatus{}, &models.SavedJob{})

	db.AutoMigrate(&models.Admin{}, &models.User{}, &models.UserSensitiveData{}, &models.Company{}, &models.CompanySensitiveData{}, &models.Job{}, &models.JobCategory{}, &models.JobStatus{}, &models.SavedJob{})

	models.AddForeignKeys(db)

	// inserting sample data

	// admin
	admin := models.Admin{
		AdminName:     "admin1",
		AdminEmail:    "admin1@beautifulcode.in",
		AdminPassword: os.Getenv("admin1_password"),
	}
	db.Save(&admin)

	// user
	mk_password := os.Getenv("mk_password")
	userSensitiveData := models.UserSensitiveData{
		UserEmail:    "mk@beautifulcode.in",
		UserPassword: mk_password,
	}
	db.Save(&userSensitiveData)
	// db.Delete(&models.UserSensitiveData{})
	// db.Delete(&models.User{})

	var userEmail models.UserSensitiveData
	db.Model(&models.UserSensitiveData{}).Select([]string{"user_email"}).First(&userEmail)

	user := models.User{
		UserName:    "mk",
		UserEmail:   userEmail.UserEmail,
		UserImage:   "https://cdn-icons-png.flaticon.com/128/552/552721.png",
		UserDoB:     time.Date(2002, time.January, 15, 0, 0, 0, 0, time.Local),
		PhoneNumber: "9999999999",
		Description: "description",
		Skills:      "skill1, skill2, skill3",
	}
	db.Create(&user)

	// company
	companySenstiveData := models.CompanySensitiveData{
		CompanyEmail:    "ops@beautifulcode.in",
		CompanyPassword: os.Getenv("company_password"),
	}
	db.Create(&companySenstiveData)

	company := models.Company{
		CompanyName:  "BeautifulCode LLP",
		CompanyEmail: companySenstiveData.CompanyEmail,
		CompanyImage: "",
		Description:  "description of company",
	}
	db.Save(&company)

	// adding job categories
	jobCategory := models.JobCategory{
		CategoryName: "Software Engineer",
	}
	db.Create(&jobCategory)
	jobCategory2 := models.JobCategory{
		CategoryName: "Human Resources",
	}
	db.Create(&jobCategory2)

	var hrCateogryId, sweCategoryId models.JobCategory

	db.Select([]string{"id"}).Where("category_name = ?", "Human Resources").Find(&hrCateogryId)
	db.Select([]string{"id"}).Where("category_name = ?", "Software Engineer").Find(&sweCategoryId)

	var bcId models.Company
	db.Select([]string{"id"}).Where("company_name = ?", "BeautifulCode LLP").Find(&bcId)

	// adding jobs
	job1 := models.Job{
		JobTitle:       "Fullstack Developer",
		JobDescription: "Job description",
		Salary:         "50000",
		CompanyID:      bcId.ID,
		CategoryID:     sweCategoryId.ID,
		// Applicants:     []*models.User{&user},
	}
	db.Create(&job1)

	job2 := models.Job{
		JobTitle:       "HR Executive",
		JobDescription: "Job description",
		Salary:         "50000",
		CompanyID:      bcId.ID,
		CategoryID:     hrCateogryId.ID,
	}
	db.Create(&job2)

	// mocking users applying to jobs
	var fsd, hr models.Job
	db.Select([]string{"id"}).Where("job_title like (?)", "Fullstack Developer").Find(&fsd)
	db.Select([]string{"id"}).Where("job_title like (?)", "HR Executive").Find(&hr)

	var u models.User
	db.Select([]string{"id"}).Where("user_name = ?", "mk").Find(&u)

	jobStatus := models.JobStatus{
		JobID:             fsd.ID,
		UserID:            u.ID,
		ApplicationStatus: models.ApplicationStatus("pending"),
	}
	db.Create(&jobStatus)

	// mocking user saving jobs
	savedJob := models.SavedJob{
		JobID:  fsd.ID,
		UserID: u.ID,
	}
	db.Create(&savedJob)

	savedJob2 := models.SavedJob{
		JobID:  hr.ID,
		UserID: u.ID,
	}
	db.Create(&savedJob2)

	return db
}
