package cmd

import (
	"errors"
	"os"

	"github.com/berlim/bridgex/src/core/models"
	"github.com/berlim/bridgex/src/core/usecases/db/create"
)

func SwitchCmd(configData models.ConfigData) error {
	var err error
	argsWithProg := os.Args[1:]
	switch argsWithProg[0] {
	case "db":
		err = switchDB(argsWithProg[1], configData)
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
		err = nil
	case "destroy":
		err = nil
	default:
		err = errors.New("this option dont exist")
	}

	return err
}
