package migrate

import (
	"fmt"
)

const SCHEMA_MIGRATION_TABLE_NAME = "schema_migrations"

func (t *dbMigrate) includeMigrationInSchema() error {
	var err error

	err = t.createSchemaMigrationTable()
	if err != nil {
		return err
	}

	return nil
}

func (t *dbMigrate) createSchemaMigrationTable() error {
	var err error

	hasTable := t.conn.HasTable(SCHEMA_MIGRATION_TABLE_NAME)

	if !hasTable {
		t.conn.Exec(schema_migration_create())

		hasTable = t.conn.HasTable(SCHEMA_MIGRATION_TABLE_NAME)
		if !hasTable {
			err = fmt.Errorf("Error in schema_migrations table generate.")
			return err
		}
	}

	return nil
}

func schema_migration_create() string {
	return fmt.Sprintf("CREATE TABLE %v ( version VARCHAR(20) NOT NULL )", SCHEMA_MIGRATION_TABLE_NAME)
}
