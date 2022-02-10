package database

import (
	"github.com/goethesum/restAPI/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and createss our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
