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

// GetAll gets all the items in the dataBase
func (s *dataBase) GetAll(data interface{}) error {
	err := s.connection.Find(data).Error
	if err != nil {
		return fmt.Errorf("failed to get all items from database: %v", err)
	}
	return nil
}

// GetByID gets an element from database using the ItemID
func (s *dataBase) GetByID(data interface{}, itemID string) error {
	err := s.connection.Where("item_id = ?", itemID).First(data).Error
	if err != nil {
		return fmt.Errorf("failed to get data from database: %v", err)
	}
	return nil
}

// CheckIfExist checks if the item exists in the database by looking for it by its ItemID
func (s *dataBase) CheckIfExist(data interface{}, itemID string) bool {
	err := s.connection.Where("item_id = ?", itemID).First(data).Error
	return err == nil
}

// Update an item in database
func (s *dataBase) Update(data interface{}, itemID string) error {
	err := s.connection.Where("item_id = ?", itemID).Updates(data).Error
	if err != nil {
		return fmt.Errorf("failed to update item in database: %v", err)
	}
	return nil
}
