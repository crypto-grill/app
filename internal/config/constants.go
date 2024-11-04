package config

import (
	"fmt"
	"strings"
)

// Hardcode other users
const (
	user1 = "http://server1:8080"
	user2 = "http://server2:8081"
	user3 = "http://server3:8082"
	user4 = "http://server4:8083"
)

var users = []string{user1, user2, user3, user4}

// GetUsersWithoutPort filters out users with a specific port
func GetUsersWithoutPort(port uint) []string {
	var filteredUsers []string
	portSuffix := fmt.Sprintf(":%d", port)

	for _, user := range users {
		// Check if the URL ends with the given port suffix
		if !strings.HasSuffix(user, portSuffix) {
			filteredUsers = append(filteredUsers, user)
		}
	}
	return filteredUsers
}
