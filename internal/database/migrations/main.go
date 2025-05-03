package migrations

import (
	"fmt"
	"github.com/mohar9h/go-graphql-ws-api/internal/database"
	"github.com/mohar9h/go-graphql-ws-api/internal/domain/group"
	"github.com/mohar9h/go-graphql-ws-api/internal/domain/permission"
	"github.com/mohar9h/go-graphql-ws-api/internal/domain/user"
	"log"
)

func RunMigrations() {
	db := database.GetDB()
	if db == nil {
		log.Fatal("❌ DB not initialized. Did you call database.InitPostgres()?")
	}

	models := []any{
		&user.User{},
		&group.Group{},
		&permission.Permission{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("✅ Migrations completed successfully!")
}
