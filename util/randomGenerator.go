package util

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// List of possible username characters
	usernameChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	// List of possible domain names
	domainNames = []string{"example.com", "example.net", "example.org", "gmail.com", "yahoo.com", "hotmail.com"}
)

func GetRandomUsername(length int) string {
	username := make([]rune, length)
	for i := range username {
		username[i] = usernameChars[rand.Intn(len(usernameChars))]
	}
	return string(username)
}

func GetRandomEmail(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random username
	username := GetRandomUsername(length) // You can specify the desired length

	// Select a random domain name
	domain := domainNames[rand.Intn(len(domainNames))]

	// Create the email address
	email := fmt.Sprintf("%s@%s", username, domain)

	return email
}

func GetRandomAmount(min, max int) int {
	// Generate a random number within the specified range
	return rand.Intn(max-min+1) + min
}
