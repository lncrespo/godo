package godo

import (
	"flag"
	"os"
)

var addCommand *flag.FlagSet
var listCommand *flag.FlagSet

type addCommandFlags struct {
	title       *string
	description *string
	priority    *int
	project     string
}

type listCommandFlags struct {
	showProjects *bool
	project      string
}

var addFlags addCommandFlags
var listFlags listCommandFlags

func init() {
	if len(os.Args) == 1 {
		FatalWithUsage("Missing subcommand.")
	}

	addCommand = flag.NewFlagSet("add", flag.ExitOnError)

	addFlags = addCommandFlags{}
	addFlags.title = addCommand.String("title", "", "")
	addCommand.StringVar(addFlags.title, "t", "", "")

	addFlags.description = addCommand.String("description", "", "")
	addCommand.StringVar(addFlags.description, "d", "", "")

	addFlags.priority = addCommand.Int("priority", 9, "")
	addCommand.IntVar(addFlags.priority, "p", 9, "")

	listCommand = flag.NewFlagSet("list", flag.ExitOnError)

	listFlags = listCommandFlags{}
	listFlags.showProjects = listCommand.Bool("projects", false, "")
	listCommand.BoolVar(listFlags.showProjects, "p", false, "")

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
		addFlags.project = addCommand.Arg(0)

		add(addFlags)
		break
	case "list":
		listCommand.Parse(os.Args[2:])
		listFlags.project = listCommand.Arg(0)

		list(listFlags)
		break
	}
}
