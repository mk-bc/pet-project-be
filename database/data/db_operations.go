package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/mk-bc/pet-project-be/models"
)

type Database interface {
	FetchAllJobs() ([]*models.Job, error)
	FetchCompanyData(companyID uint32) (*models.Company, error)
	FetchUserData(userID uint32) (*models.User, error)
	FetchJobCategories() ([]*models.JobCategory, error)
	FetchJobsByCompanyID(companyID uint32) ([]*models.Job, error)
	FetchJobsByCategoryID(categoryID uint32) ([]*models.Job, error)
	FetchCompanyJobsByCategory(companyID uint32, categoryID uint32) ([]*models.Job, error)
}

type DBClient struct {
	Db *gorm.DB
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
