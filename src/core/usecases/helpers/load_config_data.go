package helpers

import (
	"fmt"
	"io/ioutil"

	"github.com/berlim/bridgex/src/core/models"
	"gopkg.in/yaml.v3"
)

type connectionData struct {
	Host     string `yaml:"host"`
	Driver   string `yaml:"driver"`
	DbName   string `yaml:"database_name"`
	UserName string `yaml:"user_name"`
	DbPort   string `yaml:"port"`
	DbPass   string `yaml:"password"`
	SslMode  string `yaml:"ssl_mode"`
}

type path struct {
	MigrationPath  string `yaml:"migration_path"`
	ConnectionPath string `yaml:"connection_path"`
}

type dbConf struct {
	Development connectionData `yaml:"development"`
	Test        connectionData `yaml:"test"`
	Production  connectionData `yaml:"production"`
}

func LoadData(env string) (models.ConfigData, error) {

	var data models.ConfigData
	var paths path

	yamlPath, err := ioutil.ReadFile("bridgex.yml")

	ymlStr := string(yamlPath)

	err = yaml.Unmarshal(yamlPath, &paths)

	fmt.Printf("File: %v \n", ymlStr)
	fmt.Printf("Migration Path: %v \n", paths.MigrationPath)

	return data, err
}

func setConfigData(data connectionData, paths path, env string) models.ConfigData {
	configData := models.NewConfigData(
		paths.MigrationPath,
		data.Driver,
		env,
		data.DbName,
		data.Host,
		data.UserName,
		data.DbPass,
		data.DbPort,
		data.SslMode,
	)

	return configData
}

func getConnectionDataByEnv(env string, _envDB dbConf) connectionData {
	switch env {
	case "development":
		return _envDB.Development
	case "test":
		return _envDB.Test
	case "production":
		return _envDB.Production
	default:
		return _envDB.Development
	}

}
