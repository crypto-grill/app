package config

import "fmt"

// Hardcode other users, but can do with https://github.com/libp2p/go-libp2p
const (
	user1 = "http://localhost:8080"
	user2 = "http://localhost:8081"
	user3 = "http://localhost:8082"
	user4 = "http://localhost:8083"
)

var users = []string{user1, user2, user3, user4}

func GetUsersWithoutPort(port uint) []string {
	var filteredUsers []string
	addressToExclude := fmt.Sprintf("http://localhost:%d", port)

	for _, user := range users {
		if user != addressToExclude {
			filteredUsers = append(filteredUsers, user)
		}
	}
	return filteredUsers
}
