package migrations

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mohar9h/go-graphql-ws-api/internal/database"
)

func RunMigrations() {

	db := database.GetDB()

	tables, err := loadModels()
	if err != nil {
		log.Fatalf("Failed to load models: %v", err)
	}

	err = db.Migrator().CreateTable(tables...)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migrations completed successfully!")
}

func loadModels() ([]any, error) {
	var models []any

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Construct the path to the domain directory
	domainPath := filepath.Join(wd, "internal", "domain")

	err = filepath.Walk(domainPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(info.Name(), "entity.go") {
			log.Printf("Found model file: %s", path)
		}

		return nil
	})

	return models, err
}
