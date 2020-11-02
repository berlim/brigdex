package models

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
