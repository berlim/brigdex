package helpers

import (
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
	var conf dbConf

	yamlPath, err := ioutil.ReadFile("bridgex.yml")

	err = yaml.Unmarshal(yamlPath, &paths)

	if err != nil {
		return data, err
	}

	yamlConnection, err := ioutil.ReadFile(paths.ConnectionPath)

	err = yaml.Unmarshal(yamlConnection, &conf)

	if err != nil {
		return data, err
	}

	cnxData := getConnectionDataByEnv(env, conf)

	data = setConfigData(cnxData, paths, env)

	return data, nil
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
