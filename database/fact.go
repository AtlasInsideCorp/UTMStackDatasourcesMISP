package database

import (
	"fmt"
	"path/filepath"
)

// NewDBInstance returns a new instance of the database.
// The type of database (attributes) and the path where the database will be created must be passed as parameters.
func NewDBInstance(dbType, path string) (ISqlite, error) {
	if dbType == "attributes" {
		dbPath := filepath.Join(path, "attributes.db")
		return newAttributeDBInstance(dbPath)
	}
	return nil, fmt.Errorf("wrong database type passed")
}
