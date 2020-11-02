package create

import (
	"fmt"
	"os"
)

func createMigrationFiles(migrationPath string) error {
	var err error

	upFile := fmt.Sprintf("%v/up.sql", migrationPath)
	downFile := fmt.Sprintf("%v/down.sql", migrationPath)

	_, err = os.Create(upFile)
	if err != nil {
		return err
	}

	_, err = os.Create(downFile)
	if err != nil {
		return err
	}

	return nil
}
