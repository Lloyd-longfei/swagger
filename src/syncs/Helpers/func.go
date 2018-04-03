package Helpers

import (
	"os"
)

/**
	return current path
 */
func CurrentPath() string {
	path, _ := os.Getwd()
	return path
}

/**
	check for errors
 */

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
