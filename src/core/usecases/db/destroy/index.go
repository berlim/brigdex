package destroy

import (
	"fmt"

	"github.com/berlim/bridgex/src/config/db"
	"github.com/berlim/bridgex/src/core/models"
	"github.com/berlim/bridgex/src/core/usecases/helpers"
	"github.com/jinzhu/gorm"
)

type dbDrop struct {
	configData models.ConfigData
	conn       *gorm.DB
	dbName     string
}

func NewDbDrop(configData models.ConfigData) dbDrop {
	return dbDrop{
		configData: configData,
		dbName:     configData.GetDataBaseName(),
	}
}

func (t *dbDrop) Call() error {
	var err error

	t.dbName, err = helpers.RemoveSpecialCharacters(t.dbName)
	sql := fmt.Sprintf("DROP DATABASE %v;", t.dbName)

	openCnx, driver := t.configData.GetOpenConnectionBASE()
	t.conn = db.ConnectDB(openCnx, driver)

	err = t.conn.Exec(sql).Error

	if err != nil {
		return err
	}

	fmt.Printf("Database %v destroyed! \n", t.dbName)

	return nil
}
