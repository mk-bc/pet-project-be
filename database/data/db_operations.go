package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	// "github.com/mk-bc/pet-project-be/database/models"
)

type Database interface {
	// fetch

}

type DBClient struct {
	db *gorm.DB
}
