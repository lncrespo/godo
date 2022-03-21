package godo

import (
	"bufio"
	"log"
	"os"

	"github.com/lncrespo/godo/src/dbal"
)

func reset() {
	confirm, err := promptConfirmation()

	if !confirm {
		os.Stdout.WriteString("Did not receive confirmation, aborting\n")

		if err != nil {
			os.Stdout.WriteString(" due to error: " + err.Error())
		}

		return
	}

	err = dbal.TruncateTodos()

	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.WriteString("Successfully cleared all projects & todos\n")
}

func promptConfirmation() (bool, error) {
	os.Stdout.WriteString(
		"Are you sure you want to remove ALL projects & todos? This action cannot be reversed!\n")

	os.Stdout.WriteString("To confirm, type \"yes\" and press enter: ")

	reader := bufio.NewReader(os.Stdin)
	confirmText, err := reader.ReadString('\n')

	if err != nil {
		return false, err
	}

	if confirmText == "yes\n" {
		return true, nil
	}

	return false, nil
}
