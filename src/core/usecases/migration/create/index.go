package create

import (
	"fmt"

	"github.com/berlim/bridgex/src/config/db"
	"github.com/berlim/bridgex/src/core/usecases/helpers"

	"github.com/berlim/bridgex/src/core/models"
	"github.com/jinzhu/gorm"
)

type migrationCreate struct {
	configData    models.ConfigData
	conn          *gorm.DB
	dbName        string
	migrationName string
	prefix        string
}

func NewMigrationCreate(migrationName string, configData models.ConfigData) migrationCreate {
	prefix := helpers.GetDateTimeStr()
	migrationName = fmt.Sprintf("%v_%v", prefix, migrationName)

	openCnx, driver := configData.GetOpenConnection()
	conn := db.ConnectDB(openCnx, driver)

	return migrationCreate{
		configData:    configData,
		dbName:        configData.GetDataBaseName(),
		migrationName: migrationName,
		prefix:        prefix,
		conn:          conn,
	}
}

func (t *migrationCreate) Call() error {
	var err error

	migrationsPath := t.configData.GetMigrationsPath()

	err = verifyMigrationsPath(migrationsPath)
	if err != nil {
		return err
	}

	err = createMigrationDir(migrationsPath, t.migrationName)
	if err != nil {
		return err
	}

	migrationPath := fmt.Sprintf("%v/%v/", migrationsPath, t.migrationName)
	err = createMigrationFiles(migrationPath)
	if err != nil {
		return err
	}

	return nil
}
