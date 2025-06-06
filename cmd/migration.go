package main

import (
	"fmt"

	shorturl "example.com/go-shorter/internals/short_url"
)

func migrateDb() {
	fmt.Println("Migrating database tables...")

	shorturl.Migrate()

	fmt.Println("Migration successful...")
}
