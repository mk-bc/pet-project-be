package data

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/mk-bc/pet-project-be/models"
)

func newMockDBClient() (*gorm.DB, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New()
	gormDB, err := gorm.Open("postgres", mockDb)
	if err != nil {
		panic("Unexpected error while connection to mock sql: ")
	}
	return gormDB, mock
}

func TestCreateProduct(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	dbClinet := DBClient{
		Db: db,
	}

	t.Run("Create New Job", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "jobs"(.+)`).WillReturnRows(sqlmock.NewRows([]string{"id", "job_title", "job_description", "salary", "company_id", "category_id"}).AddRow(1, "Mock Job", "Mock Description", "50000", 1, 1))
		mock.ExpectCommit()

		jobs, err := dbClinet.CreateNewJob(models.Job{
			JobTitle:       "Mock Job",
			JobDescription: "Mock description",
			Salary:         "50000",
			CompanyID:      1,
			CategoryID:     1,
		})
		log.Println(err, jobs)
		// log.Printf("%v %v %v", jobs.ID, jobs.JobTitle, jobs.Salary)

	})
}

func TestRegisterUser(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	dbClient := DBClient{
		Db: db,
	}

	t.Run("Register User", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "sensitive_data" (.+)`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := dbClient.RegisterUser(
			&models.SensitiveData{
				Email:    "mock-user@email.com",
				Password: "mock",
				Role:     "applicant",
			}, &models.User{
				UserName:    "mock-name",
				Description: "mock-description",
			},
		)

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("%v", err)
		}
	})
}

func TestFetchJobs(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	dbClient := DBClient{
		Db: db,
	}

	t.Run("Fetch all jobs", func(t *testing.T) {
		// mock.ExpectBegin()
		mockr := sqlmock.NewRows([]string{"id", "job_title"}).
			AddRow(1, "SWE").
			AddRow(2, "HR").
			AddRow(3, "PM")
		mock.ExpectQuery(`SELECT *.`).
			// WillReturnError(fmt.Errorf("%v", "error srinu"))
			WillReturnRows(mockr)
		mock.ExpectCommit()

		got, err := dbClient.FetchAllJobs()

		t.Error(got, err)
		if got == nil {
			t.Error(got)
		}
		log.Println(err, got)
	})
}
