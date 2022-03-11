package godo

import (
	"flag"
	"os"
)

var addCommand *flag.FlagSet
var listCommand *flag.FlagSet
var addTitle *string
var addDescription *string
var addPriority *int
var project string

func init() {
	if len(os.Args) == 1 {
		FatalWithUsage("Missing subcommand.")
	}

	addCommand = flag.NewFlagSet("add", flag.ExitOnError)

	addTitle = addCommand.String("title", "", "")
	addCommand.StringVar(addTitle, "t", "", "")

	addDescription = addCommand.String("description", "", "")
	addCommand.StringVar(addDescription, "d", "", "")

	addPriority = addCommand.Int("priority", 9, "")
	addCommand.IntVar(addPriority, "p", 9, "")

	listCommand = flag.NewFlagSet("list", flag.ExitOnError)

	addCommand.Usage = func() {
		ExitWithUsage()
	}

	listCommand.Usage = func() {
		ExitWithUsage()
	}

	flag.Usage = func() {
		ExitWithUsage()
	}

	flag.Parse()
}

func ParseSubcommands() {
	subcommand := os.Args[1]

	switch subcommand {
	case "add":
		addCommand.Parse(os.Args[2:])
		project = addCommand.Arg(0)

		add(*addTitle, *addDescription, *addPriority, project)
		break
	case "list":
		listCommand.Parse(os.Args[2:])
		project = listCommand.Arg(0)

		list(project)
		break
	}
}
