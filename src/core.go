package godo

import (
	"flag"
	"fmt"
	"os"
)

var addTitle *string
var addDescription *string
var addPriority *int
var addProject string

func init() {
	if len(os.Args) == 1 {
		FatalWithUsage("Missing subcommand.")
	}

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	addTitle = addCommand.String("title", "", "")
	addCommand.StringVar(addTitle, "t", "", "")

	addDescription = addCommand.String("description", "", "")
	addCommand.StringVar(addDescription, "d", "", "")

	addPriority = addCommand.Int("priority", 9, "")
	addCommand.IntVar(addPriority, "p", 9, "")

	addCommand.Usage = func() {
		ExitWithUsage()
	}

	flag.Usage = func() {
		ExitWithUsage()
	}

	addCommand.Parse(os.Args[2:])
	flag.Parse()

	addProject = addCommand.Arg(0)
}

func ParseSubcommands() {
	subcommand := os.Args[1]

	fmt.Println(subcommand)

	initializeStorage()

	switch subcommand {
	case "add":
		add(*addTitle, *addDescription, *addPriority, addProject)
		break
	}
}

func printUsage() {
}
