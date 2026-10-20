package analytics

import (
	"context"
	"fmt"
	"xinsight/internals/xapi"
)

func AnalyzeUser(ctx context.Context, client *xapi.Client, username string) error {
	fmt.Println("Starting analytics...")
	return nil
}