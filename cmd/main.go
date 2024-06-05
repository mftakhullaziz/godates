package main

import (
	"context"
	"godating-dealls/config"
)

func main() {
	// Init context before run application
	ctx := context.Background()
	// Ensure to close the database connection when the application exits
	defer config.CloseDBConnection()
	// Create of the database connection
	db := config.CreateDBConnection(ctx)

	defer db.Close()
}
