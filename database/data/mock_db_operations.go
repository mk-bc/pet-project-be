// Code generated by MockGen. DO NOT EDIT.
// Source: db_operations.go

// Package data is a generated GoMock package.
package data

import (
	reflect "reflect"

	models "github.com/mk-bc/pet-project-be/models"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CheckAppliedJobs mocks base method.
func (m *MockDatabase) CheckAppliedJobs(arg0 uint32) ([]*models.JobStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAppliedJobs", arg0)
	ret0, _ := ret[0].([]*models.JobStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAppliedJobs indicates an expected call of CheckAppliedJobs.
func (mr *MockDatabaseMockRecorder) CheckAppliedJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAppliedJobs", reflect.TypeOf((*MockDatabase)(nil).CheckAppliedJobs), arg0)
}

// CheckSavedJobs mocks base method.
func (m *MockDatabase) CheckSavedJobs(arg0 uint32) ([]*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSavedJobs", arg0)
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSavedJobs indicates an expected call of CheckSavedJobs.
func (mr *MockDatabaseMockRecorder) CheckSavedJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSavedJobs", reflect.TypeOf((*MockDatabase)(nil).CheckSavedJobs), arg0)
}

// CreateNewJob mocks base method.
func (m *MockDatabase) CreateNewJob(arg0 models.Job) (*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewJob", arg0)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewJob indicates an expected call of CreateNewJob.
func (mr *MockDatabaseMockRecorder) CreateNewJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewJob", reflect.TypeOf((*MockDatabase)(nil).CreateNewJob), arg0)
}

// CreateNewJobCategory mocks base method.
func (m *MockDatabase) CreateNewJobCategory(arg0 models.JobCategory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewJobCategory", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewJobCategory indicates an expected call of CreateNewJobCategory.
func (mr *MockDatabaseMockRecorder) CreateNewJobCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewJobCategory", reflect.TypeOf((*MockDatabase)(nil).CreateNewJobCategory), arg0)
}

// DeleteCompany mocks base method.
func (m *MockDatabase) DeleteCompany(companyID uint32, companyEmail string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", companyID, companyEmail)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockDatabaseMockRecorder) DeleteCompany(companyID, companyEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockDatabase)(nil).DeleteCompany), companyID, companyEmail)
}

// DeleteJob mocks base method.
func (m *MockDatabase) DeleteJob(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJob indicates an expected call of DeleteJob.
func (mr *MockDatabaseMockRecorder) DeleteJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJob", reflect.TypeOf((*MockDatabase)(nil).DeleteJob), arg0)
}

// DeleteUser mocks base method.
func (m *MockDatabase) DeleteUser(arg0 uint32, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockDatabaseMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockDatabase)(nil).DeleteUser), arg0, arg1)
}

// FetchAllJobs mocks base method.
func (m *MockDatabase) FetchAllJobs() ([]*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllJobs")
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllJobs indicates an expected call of FetchAllJobs.
func (mr *MockDatabaseMockRecorder) FetchAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllJobs", reflect.TypeOf((*MockDatabase)(nil).FetchAllJobs))
}

// FetchApplicantsByJobID mocks base method.
func (m *MockDatabase) FetchApplicantsByJobID(arg0 uint32) ([]*models.JobStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchApplicantsByJobID", arg0)
	ret0, _ := ret[0].([]*models.JobStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchApplicantsByJobID indicates an expected call of FetchApplicantsByJobID.
func (mr *MockDatabaseMockRecorder) FetchApplicantsByJobID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchApplicantsByJobID", reflect.TypeOf((*MockDatabase)(nil).FetchApplicantsByJobID), arg0)
}

// FetchCompanyData mocks base method.
func (m *MockDatabase) FetchCompanyData(arg0 uint32) (*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCompanyData", arg0)
	ret0, _ := ret[0].(*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCompanyData indicates an expected call of FetchCompanyData.
func (mr *MockDatabaseMockRecorder) FetchCompanyData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCompanyData", reflect.TypeOf((*MockDatabase)(nil).FetchCompanyData), arg0)
}

// FetchCompanyDataByEmail mocks base method.
func (m *MockDatabase) FetchCompanyDataByEmail(arg0 string) (*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCompanyDataByEmail", arg0)
	ret0, _ := ret[0].(*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCompanyDataByEmail indicates an expected call of FetchCompanyDataByEmail.
func (mr *MockDatabaseMockRecorder) FetchCompanyDataByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCompanyDataByEmail", reflect.TypeOf((*MockDatabase)(nil).FetchCompanyDataByEmail), arg0)
}

// FetchCompanyJobsByCategory mocks base method.
func (m *MockDatabase) FetchCompanyJobsByCategory(arg0, arg1 uint32) ([]*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCompanyJobsByCategory", arg0, arg1)
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCompanyJobsByCategory indicates an expected call of FetchCompanyJobsByCategory.
func (mr *MockDatabaseMockRecorder) FetchCompanyJobsByCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCompanyJobsByCategory", reflect.TypeOf((*MockDatabase)(nil).FetchCompanyJobsByCategory), arg0, arg1)
}

// FetchJobCategories mocks base method.
func (m *MockDatabase) FetchJobCategories() ([]*models.JobCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobCategories")
	ret0, _ := ret[0].([]*models.JobCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobCategories indicates an expected call of FetchJobCategories.
func (mr *MockDatabaseMockRecorder) FetchJobCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobCategories", reflect.TypeOf((*MockDatabase)(nil).FetchJobCategories))
}

// FetchJobData mocks base method.
func (m *MockDatabase) FetchJobData(arg0 uint32) (*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobData", arg0)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobData indicates an expected call of FetchJobData.
func (mr *MockDatabaseMockRecorder) FetchJobData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobData", reflect.TypeOf((*MockDatabase)(nil).FetchJobData), arg0)
}

// FetchJobsByCategoryID mocks base method.
func (m *MockDatabase) FetchJobsByCategoryID(arg0 uint32) ([]*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobsByCategoryID", arg0)
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobsByCategoryID indicates an expected call of FetchJobsByCategoryID.
func (mr *MockDatabaseMockRecorder) FetchJobsByCategoryID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobsByCategoryID", reflect.TypeOf((*MockDatabase)(nil).FetchJobsByCategoryID), arg0)
}

// FetchJobsByCompanyID mocks base method.
func (m *MockDatabase) FetchJobsByCompanyID(arg0 uint32) ([]*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobsByCompanyID", arg0)
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobsByCompanyID indicates an expected call of FetchJobsByCompanyID.
func (mr *MockDatabaseMockRecorder) FetchJobsByCompanyID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobsByCompanyID", reflect.TypeOf((*MockDatabase)(nil).FetchJobsByCompanyID), arg0)
}

// FetchUserData mocks base method.
func (m *MockDatabase) FetchUserData(arg0 uint32) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserData", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserData indicates an expected call of FetchUserData.
func (mr *MockDatabaseMockRecorder) FetchUserData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserData", reflect.TypeOf((*MockDatabase)(nil).FetchUserData), arg0)
}

// FetchUserDataByEmail mocks base method.
func (m *MockDatabase) FetchUserDataByEmail(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserDataByEmail", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserDataByEmail indicates an expected call of FetchUserDataByEmail.
func (mr *MockDatabaseMockRecorder) FetchUserDataByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserDataByEmail", reflect.TypeOf((*MockDatabase)(nil).FetchUserDataByEmail), arg0)
}

// Login mocks base method.
func (m *MockDatabase) Login(email string) (*models.SensitiveData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email)
	ret0, _ := ret[0].(*models.SensitiveData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockDatabaseMockRecorder) Login(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockDatabase)(nil).Login), email)
}

// ModifyApplicationStatus mocks base method.
func (m *MockDatabase) ModifyApplicationStatus(arg0 models.JobStatus) (*models.JobStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyApplicationStatus", arg0)
	ret0, _ := ret[0].(*models.JobStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyApplicationStatus indicates an expected call of ModifyApplicationStatus.
func (mr *MockDatabaseMockRecorder) ModifyApplicationStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyApplicationStatus", reflect.TypeOf((*MockDatabase)(nil).ModifyApplicationStatus), arg0)
}

// RegisterCompany mocks base method.
func (m *MockDatabase) RegisterCompany(arg0 *models.SensitiveData, arg1 *models.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCompany", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterCompany indicates an expected call of RegisterCompany.
func (mr *MockDatabaseMockRecorder) RegisterCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCompany", reflect.TypeOf((*MockDatabase)(nil).RegisterCompany), arg0, arg1)
}

// RegisterUser mocks base method.
func (m *MockDatabase) RegisterUser(arg0 *models.SensitiveData, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockDatabaseMockRecorder) RegisterUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockDatabase)(nil).RegisterUser), arg0, arg1)
}

// UpdateCompanyData mocks base method.
func (m *MockDatabase) UpdateCompanyData(arg0 models.Company, arg1 models.SensitiveData) (*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompanyData", arg0, arg1)
	ret0, _ := ret[0].(*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompanyData indicates an expected call of UpdateCompanyData.
func (mr *MockDatabaseMockRecorder) UpdateCompanyData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompanyData", reflect.TypeOf((*MockDatabase)(nil).UpdateCompanyData), arg0, arg1)
}

// UpdateJobData mocks base method.
func (m *MockDatabase) UpdateJobData(arg0 models.Job, arg1 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJobData indicates an expected call of UpdateJobData.
func (mr *MockDatabaseMockRecorder) UpdateJobData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobData", reflect.TypeOf((*MockDatabase)(nil).UpdateJobData), arg0, arg1)
}

// UpdateUserData mocks base method.
func (m *MockDatabase) UpdateUserData(arg0 models.User, arg1 models.SensitiveData) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserData", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserData indicates an expected call of UpdateUserData.
func (mr *MockDatabaseMockRecorder) UpdateUserData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserData", reflect.TypeOf((*MockDatabase)(nil).UpdateUserData), arg0, arg1)
}

// UserJobApplication mocks base method.
func (m *MockDatabase) UserJobApplication(arg0 models.JobStatus) (*models.JobStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserJobApplication", arg0)
	ret0, _ := ret[0].(*models.JobStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserJobApplication indicates an expected call of UserJobApplication.
func (mr *MockDatabaseMockRecorder) UserJobApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserJobApplication", reflect.TypeOf((*MockDatabase)(nil).UserJobApplication), arg0)
}

// UserRemoveSavedJob mocks base method.
func (m *MockDatabase) UserRemoveSavedJob(arg0 models.SavedJob) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRemoveSavedJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserRemoveSavedJob indicates an expected call of UserRemoveSavedJob.
func (mr *MockDatabaseMockRecorder) UserRemoveSavedJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRemoveSavedJob", reflect.TypeOf((*MockDatabase)(nil).UserRemoveSavedJob), arg0)
}

// UserSavedJob mocks base method.
func (m *MockDatabase) UserSavedJob(arg0 models.SavedJob) (*models.SavedJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSavedJob", arg0)
	ret0, _ := ret[0].(*models.SavedJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSavedJob indicates an expected call of UserSavedJob.
func (mr *MockDatabaseMockRecorder) UserSavedJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSavedJob", reflect.TypeOf((*MockDatabase)(nil).UserSavedJob), arg0)
}
