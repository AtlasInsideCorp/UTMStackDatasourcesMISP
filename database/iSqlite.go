package database

type ISqlite interface {
	Migrate(data interface{}) error                                              // Migrate create a new table in the dataBase
	AddNew(data interface{}) error                                               // AddNew adds a new item to the dataBase
	GetAll(data interface{}) error                                               // GetAll gets all the items in the dataBase
	GetByID(data interface{}, itemID string) error                               // GetByID gets an element from database using the ItemID
	CheckIfExist(data interface{}, itemID string) bool                           // CheckIfExist checks if the item exists in the database by looking for it by its ItemID
	Update(data interface{}, itemID string) error                                // Update an item in database
	UpdateOnlyField(data, newValue interface{}, itemID, columnName string) error // UpdateOnlyField updates only one field of an item in the database
	Delete(data interface{}, itemID string) error                                // Delete an element in database
	Close() error                                                                // Close the dataBase
}
