package consumer

import "fmt"

// User is a user of the service
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (user *User) String() string {
	return fmt.Sprintf("Users Name:  %s\nUsers Email: %s", user.Name, user.Email)
}
