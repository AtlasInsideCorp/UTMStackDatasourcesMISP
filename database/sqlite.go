package database

import (
	"fmt"

	"gorm.io/gorm"
)

type dataBase struct {
	connection *gorm.DB
}

// Migrate create a new table in the dataBase
func (s *dataBase) Migrate(data interface{}) error {
	err := s.connection.AutoMigrate(&data)
	if err != nil {
		return fmt.Errorf("failed to create new table in database: %v", err)
	}
	return nil
}

// Close the dataBase
func (s *dataBase) Close() error {
	sqlDB, err := s.connection.DB()
	if err != nil {
		return fmt.Errorf("failed to close database: %v", err)
	}
	sqlDB.Close()
	return nil
}

// AddNew adds a new item to the dataBase
func (s *dataBase) AddNew(data interface{}) error {
	err := s.connection.Create(data).Error
	if err != nil {
		return fmt.Errorf("failed to add new item to database: %v", err)
	}
	return nil
}
