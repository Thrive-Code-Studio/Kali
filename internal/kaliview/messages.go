package kaliview

import (
	"fmt"
)


const (
	mainMessage string = "Welcome to Kali!"
	secondMessage string = "Select the right options for your needs"
)

func messageJoiner() string {
	return fmt.Sprintf("%s\n\n%s", mainMessage, secondMessage)
}
