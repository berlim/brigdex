package create

import (
	"fmt"
	"os"
)

func verifyMigrationsPath(migrationsPath string) error {
	var err error

	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		err = os.Mkdir(migrationsPath, os.ModePerm)
	}

	return err
}

func createMigrationDir(migrationsPath string, migrationName string) error {
	var err error

	dir := fmt.Sprintf("%v/%v", migrationsPath, migrationName)

	err = os.Mkdir(dir, os.ModePerm)

	return err
}
