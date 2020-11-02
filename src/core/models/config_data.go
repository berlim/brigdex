package models

import (
	"fmt"
)

type ConfigData struct {
	migrationPath string
	driver        string
	env           string
	dbName        string
	dbHost        string
	dbPort        string
	dbUser        string
	dbPass        string
	dbSsl         string
}

func NewConfigData(migration string, dbDriver string, env string, dbName string, dbHost string, dbUser string, dbPass string, dbPort string, dbSsl string) ConfigData {
	return ConfigData{
		migrationPath: migration,
		driver:        dbDriver,
		env:           env,
		dbName:        dbName,
		dbHost:        dbHost,
		dbPort:        dbPort,
		dbUser:        dbUser,
		dbPass:        dbPass,
		dbSsl:         dbSsl,
	}
}

func (t *ConfigData) GetDataBaseName() string {
	return t.dbName
}

func (t *ConfigData) GetOpenConnection() (string, string) {
	open := fmt.Sprintf("host=%v user=%v dbname=%v sslmode=%v password=%v port=%v",
		t.dbHost,
		t.dbUser,
		t.dbName,
		t.dbSsl,
		t.dbPass,
		t.dbPort,
	)

	return open, t.driver
}

func (t *ConfigData) GetOpenConnectionBASE() (string, string) {
	open := fmt.Sprintf("host=%v user=%v dbname=postgres sslmode=%v password=%v port=%v",
		t.dbHost,
		t.dbUser,
		t.dbSsl,
		t.dbPass,
		t.dbPort,
	)

	return open, t.driver
}
