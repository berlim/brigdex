package cmd

import (
	"errors"
	"os"

	"github.com/berlim/bridgex/src/core/models"
	"github.com/berlim/bridgex/src/core/usecases/db/create"
	"github.com/berlim/bridgex/src/core/usecases/db/destroy"
	"github.com/berlim/bridgex/src/core/usecases/db/migrate"
	_migrationCreate "github.com/berlim/bridgex/src/core/usecases/migration/create"
)

func SwitchCmd(configData models.ConfigData) error {
	var err error
	argsWithProg := os.Args[1:]
	switch argsWithProg[0] {
	case "db":
		err = switchDB(argsWithProg[1], configData)
		break
	case "migration":
		err = switchMigration(argsWithProg[1], argsWithProg[2], configData)
		break
	default:
		err = errors.New("this command dont exist")
	}
	return err
}

func switchDB(option string, configData models.ConfigData) error {
	var err error

	switch option {
	case "create":
		dbCreate := create.NewDbCreate(configData)
		err = dbCreate.Call()
		break
	case "migrate":
		dbMigrate := migrate.NewDbMigrate(configData)
		err = dbMigrate.Call()
		break
	case "destroy":
		dbDestroy := destroy.NewDbDrop(configData)
		err = dbDestroy.Call()
		break
	default:
		err = errors.New("this option dont exist")
	}

	return err
}

func switchMigration(option string, nameMigration string, configData models.ConfigData) error {
	var err error

	switch option {
	case "create":
		migrationCreate := _migrationCreate.NewMigrationCreate(nameMigration, configData)
		err = migrationCreate.Call()
		break
	case "rollback":
		dbCreate := create.NewDbCreate(configData)
		err = dbCreate.Call()
		break
	case "destroy":
		dbDestroy := destroy.NewDbDrop(configData)
		err = dbDestroy.Call()
		break
	default:
		err = errors.New("this option dont exist")
	}

	return err
}
