package init

import (
	"io/ioutil"

	"github.com/berlim/bridgex/src/core/models"
	"gopkg.in/yaml.v3"
)

type connectionData struct {
	Host     string `yaml:"host"`
	DbName   string `yaml:"database_name"`
	UserName string `yaml:"user_name"`
	DbPort   string `yaml:"port"`
	DbPass   string `yaml:"password"`
}

type path struct {
	migrationPath  string `yaml:"migration_path"`
	connectionPath string `yaml:"connection_path"`
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

	err = yaml.Unmarshal(yamlPath, &paths)

	return data, err
}
