package models

type ConfigData struct {
	migrationPath string
	driver        string
	env           string
	dbName        string
	dbHost        string
	dbUser        string
	dbPass        string
	dbSll         string
}
