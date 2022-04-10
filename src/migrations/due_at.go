package migrations

import "github.com/lncrespo/godo/src/dbal"

func migrateDueAt() error {
	statement, err := dbal.Db.Prepare(
		"ALTER TABLE `todo` ADD COLUMN `due_at` DATETIME NOT NULL DEFAULT '0001-01-01 00:00:00'")

	if err != nil {
		return err
	}

	_, err = statement.Exec()

	return err
}
