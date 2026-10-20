package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"xinsight/internals/xapi"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "analyze" {
		fmt.Println("XInsight - X Analytics CLI")
		fmt.Println("\nUsage:")
		fmt.Println("  go run . analyze -username <username>")
		return
	}

	analyzeCmd := flag.NewFlagSet("analyze", flag.ExitOnError)
	username := analyzeCmd.String(
		"username", "", "X username to analyze",
	)

	_ = analyzeCmd.Parse(os.Args[2:])

	if strings.TrimSpace(*username) == "" {
		fmt.Fprintln(os.Stderr, "Error: username is required")
		os.Exit(1)
	}

	client, err := xapi.NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	user, err := client.GetUser(context.Background(), *username)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println("\n========== XINSIGHT ==========")
	fmt.Printf("Name:       %s\n", user.Name)
	fmt.Printf("Username:   @%s\n", user.Username)
	fmt.Printf("User ID:    %s\n", user.ID)
	fmt.Printf("Followers:  %d\n", user.PublicMetrics.FollowersCount)
	fmt.Printf("Following:  %d\n", user.PublicMetrics.FollowingCount)
	fmt.Printf("Posts:      %d\n", user.PublicMetrics.TweetCount)
	fmt.Printf("Bio:        %s\n", user.Description)
	fmt.Println("==============================")
}