package utils

import "os"

// FileExist ...
func FileExist(file string) bool {

	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return true

}
