package database

import (
	"sync"
)

var (
	attributesInstance    *dataBase
	sqliteAttributeDBOnce sync.Once
)

// newAttributeDBInstance returns a single database instance of type Attributes.
// The path where the database will be created must be passed as a parameter.
func newAttributeDBInstance(path string) (ISqlite, error) {
	var err error
	sqliteAttributeDBOnce.Do(func() {
		attributesInstance, err = newDB(path)
	})
	return attributesInstance, err
}
