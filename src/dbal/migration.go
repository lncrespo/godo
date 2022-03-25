package dbal

import (
	"errors"
	"log"
)

func GetMigrationVersion() (int, error) {
	migrationVersion := -1

	if Db == nil {
		return migrationVersion, errors.New("Database connection is not established.")
	}

	statement, err := Db.Prepare(
		"SELECT * FROM `migration_version`")

	if err != nil {
		if err.Error() == "no such table: migration_version" {
			log.Println("Could not find `migration_version` table. Creating table")

			createMigrationTable()
		} else {
			return migrationVersion, err
		}
	}

	// We need to reinitialize the statement variable here, since db.Prepare() returns nil for the
	// statement if err is not nil
	statement, err = Db.Prepare(
		"SELECT * FROM `migration_version`")

	if err != nil {
		return migrationVersion, err
	}

	err = statement.QueryRow().Scan(&migrationVersion)

	return migrationVersion, err
}

func SetMigrationVersion(version int) error {
	if Db == nil {
		return errors.New("Database connection is not established.")
	}

	statement, err := Db.Prepare("UPDATE `migration_version` SET `version` = ?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(version)

	return err
}
