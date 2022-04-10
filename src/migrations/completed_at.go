package migrations

import "github.com/lncrespo/godo/src/dbal"

func migrateCompletedAt() error {
	statement, err := dbal.Db.Prepare(
		"ALTER TABLE `todo` ADD COLUMN `completed_at` DATETIME NOT NULL DEFAULT '0000-01-01 12:00:00'")

	if err != nil {
		return err
	}

	_, err = statement.Exec()

	return err
}
