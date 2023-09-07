package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/mk-bc/pet-project-be/models"
)

type Database interface {
	Login(email string) (*models.SensitiveData, error)

	RegisterCompany(*models.SensitiveData, *models.Company) error
	DeleteCompany(companyID uint32, companyEmail string) error
	UpdateCompanyData(models.Company, models.SensitiveData) (*models.Company, error)

	CreateNewJob(models.Job) (*models.Job, error)
	CreateNewJobCategory(models.JobCategory) error
	UpdateJobData(models.Job, uint32) error
	DeleteJob(uint32) error
	FetchApplicantsByJobID(uint32) ([]*models.JobStatus, error)
	ModifyApplicationStatus(models.JobStatus) (*models.JobStatus, error)

	FetchAllJobs() ([]*models.Job, error)
	FetchCompanyData(uint32) (*models.Company, error)
	FetchUserData(uint32) (*models.User, error)
	FetchJobCategories() ([]*models.JobCategory, error)
	FetchJobsByCompanyID(uint32) ([]*models.Job, error)
	FetchJobsByCategoryID(uint32) ([]*models.Job, error)
	FetchCompanyJobsByCategory(uint32, uint32) ([]*models.Job, error)

	FetchJobData(uint32) (*models.Job, error)
	FetchCompanyDataByEmail(string) (*models.Company, error) //no rpc call yet
}

type DBClient struct {
	Db *gorm.DB
}

func (db *DBClient) Login(email string) (*models.SensitiveData, error) {
	var data models.SensitiveData
	if err := db.Db.Where("email = ?", email).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (db *DBClient) FetchAllJobs() ([]*models.Job, error) {
	var jobs []*models.Job
	if err := db.Db.Model(&models.Job{}).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *DBClient) FetchCompanyData(companyID uint32) (*models.Company, error) {
	var company models.Company
	if err := db.Db.Where("id = ?", companyID).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (db *DBClient) FetchUserData(userID uint32) (*models.User, error) {
	var user models.User
	if err := db.Db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DBClient) FetchJobCategories() ([]*models.JobCategory, error) {
	var categories []*models.JobCategory
	if err := db.Db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (db *DBClient) FetchJobsByCompanyID(companyID uint32) ([]*models.Job, error) {
	var jobs []*models.Job
	if err := db.Db.Where("company_id = ?", companyID).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *DBClient) FetchJobsByCategoryID(categoryID uint32) ([]*models.Job, error) {
	var jobs []*models.Job
	if err := db.Db.Where("category_id = ?", categoryID).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *DBClient) FetchCompanyJobsByCategory(companyID uint32, categoryID uint32) ([]*models.Job, error) {
	var jobs []*models.Job
	if err := db.Db.Where("company_id = ? and category_id = ?", companyID, categoryID).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *DBClient) CreateNewJob(job models.Job) (*models.Job, error) {
	if err := db.Db.Create(&job).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (db *DBClient) RegisterCompany(credentials *models.SensitiveData, company *models.Company) error {
	if err := db.Db.Create(&credentials).Error; err != nil {
		return fmt.Errorf("Error credentials: %v", err)
	}
	if err := db.Db.Create(&company).Error; err != nil {
		return fmt.Errorf("Error inserting company: %v", err)
	}
	return nil
}

func (db *DBClient) DeleteCompany(companyID uint32, companyEmail string) error {
	// onDelete - CASCADE => record in company table is also deleted
	if err := db.Db.Where("email = ?", companyEmail).Delete(&models.SensitiveData{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DBClient) UpdateCompanyData(company models.Company, creds models.SensitiveData) (*models.Company, error) {
	if err := db.Db.Model(&models.Company{}).Where("company_email = ?", creds.Email).Updates(models.Company{
		CompanyName:  company.CompanyName,
		CompanyImage: company.CompanyImage,
		Description:  company.Description,
	}).Error; err != nil {
		return nil, err
	}
	return &company, nil
}
func (db *DBClient) CreateNewJobCategory(newCategory models.JobCategory) error {
	if err := db.Db.Save(&newCategory).Error; err != nil {
		return err
	}
	return nil
}

func (db *DBClient) UpdateJobData(job models.Job, jobID uint32) error {
	if err := db.Db.Model(&models.Job{}).Where("id = ?", jobID).Updates(models.Job{
		JobTitle:       job.JobTitle,
		JobDescription: job.JobDescription,
		Salary:         job.Salary,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DBClient) DeleteJob(jobID uint32) error {
	if err := db.Db.Debug().Where("id = ?", jobID).Delete(&models.Job{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DBClient) FetchJobData(jobID uint32) (*models.Job, error) {
	var job models.Job
	if err := db.Db.Where("id = ?", jobID).First(&job).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (db *DBClient) FetchCompanyDataByEmail(email string) (*models.Company, error) {
	var company models.Company
	if err := db.Db.Where("company_email = ?", email).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (db *DBClient) FetchApplicantsByJobID(jobID uint32) ([]*models.JobStatus, error) {
	var applicantStatuses []*models.JobStatus
	if err := db.Db.Where("job_id = ?", jobID).Find(&applicantStatuses).Error; err != nil {
		return nil, err
	}
	return applicantStatuses, nil
}

func (db *DBClient) ModifyApplicationStatus(jobStatus models.JobStatus) (*models.JobStatus, error) {
	if err := db.Db.Model(&models.JobStatus{}).Where("job_id = ? and user_id = ?", jobStatus.JobID, jobStatus.UserID).Updates(&models.JobStatus{
		ApplicationStatus: jobStatus.ApplicationStatus,
	}).Error; err != nil {
		return nil, err
	}
	return &jobStatus, nil
}
