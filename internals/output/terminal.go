package output

import (
	"fmt"
	"xinsight/internals/xapi"
)

func PrintUserDetails(user *xapi.User) {
	fmt.Printf("User: %s\n", user.Username)
	fmt.Printf("Description: %s\n", user.Description)
	fmt.Printf("Created at: %s\n", user.CreatedAt)
	fmt.Printf("Public metrics: %v\n", user.PublicMetrics)
	fmt.Printf("Verified: %t\n", user.Verified)
	fmt.Printf("Protected: %t\n", user.Protected)
}