package main

import (
	"context"
	"godating-dealls/conf"
)

func main() {
	// Init context before run application
	ctx := context.Background()
	// Ensure to close the database connection when the application exits
	defer conf.CloseDBConnection()
	// Create of the database connection
	_ = conf.CreateDBConnection(ctx)

	//defer db.Close()
}
