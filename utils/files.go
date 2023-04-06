package utils

import "os"

// CheckIfExistPath checks if a specific path exists
func CheckIfExistPath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
