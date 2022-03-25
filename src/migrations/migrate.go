package migrations

import (
	"log"
	"os"
	"strconv"

	"github.com/lncrespo/godo/src/dbal"
)

var migrations = []func() error{
	migrateCompletedAt,
}

func ExecuteMigrations() {
	currentVersion, err := dbal.GetMigrationVersion()
	executedMigrations := false

	if err != nil {
		log.Fatalln(err)
	}

	for version, migration := range migrations {
		if version < currentVersion {
			continue
		}

		err := migration()

		if err != nil {
			log.Fatalln(
				"Error while executing migration " + strconv.Itoa(version) + ": " + err.Error())

			break
		}

		executedMigrations = true
		currentVersion = version + 1
	}

	err = dbal.SetMigrationVersion(currentVersion)

	if err != nil {
		log.Fatalln(err)
	}

	if executedMigrations {
		os.Stdout.WriteString("Migrated up to version " + strconv.Itoa(currentVersion) + "\n")
	}
}
