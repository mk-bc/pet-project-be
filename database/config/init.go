package config

import (
	"log"
	"os"

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

	db.DropTableIfExists(&models.Admin{}, &models.User{}, &models.SensitiveData{}, &models.Company{}, &models.Job{}, &models.JobCategory{}, &models.JobStatus{}, &models.SavedJob{})

	db.AutoMigrate(&models.Admin{}, &models.User{}, &models.SensitiveData{}, &models.Company{}, &models.Job{}, &models.JobCategory{}, &models.JobStatus{}, &models.SavedJob{})

	models.AddForeignKeys(db)

	// inserting sample data

	// admin
	adminCredentials := models.SensitiveData{
		Email:    "admin1@bcode.in",
		Password: os.Getenv("admin1_password"),
		Role:     "admin",
	}
	db.Save(&adminCredentials)

	admin := models.Admin{
		AdminName:  "admin1",
		AdminEmail: "admin1@bcode.in",
	}
	db.Save(&admin)

	// user
	userCredentials := models.SensitiveData{
		Email:    "mk@bcode.in",
		Password: os.Getenv("mk_password"),
		Role:     "applicant",
	}
	db.Save(&userCredentials)

	var userEmail models.SensitiveData
	db.Model(&models.SensitiveData{}).Select([]string{"email"}).Where("role = ?", "applicant").First(&userEmail)

	user := models.User{
		UserName:    "mk",
		UserEmail:   userCredentials.Email,
		UserImage:   "https://cdn-icons-png.flaticon.com/128/552/552721.png",
		PhoneNumber: "9999999999",
		Description: "description",
		Skills:      "skill1, skill2, skill3",
	}
	db.Create(&user)

	// company
	companySenstiveData := models.SensitiveData{
		Email:    "ops@bcode.in",
		Password: os.Getenv("company_password"),
		Role:     "company",
	}
	db.Create(&companySenstiveData)

	company := models.Company{
		CompanyName:  "BeautifulCode LLP",
		CompanyEmail: companySenstiveData.Email,
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
