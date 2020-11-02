package migrate

import (
	"github.com/berlim/bridgex/src/config/db"
	"github.com/berlim/bridgex/src/core/models"
	"github.com/jinzhu/gorm"
)

type dbMigrate struct {
	configData models.ConfigData
	conn       *gorm.DB
	dbName     string
}

func NewDbMigrate(configData models.ConfigData) dbMigrate {

	openCnx, driver := configData.GetOpenConnection()
	conn := db.ConnectDB(openCnx, driver)

	return dbMigrate{
		configData: configData,
		dbName:     configData.GetDataBaseName(),
		conn:       conn,
	}
}

func (t *dbMigrate) Call() error {
	var err error

	err = t.includeMigrationInSchema()

	if err != nil {
		return err
	}

	return nil
}
