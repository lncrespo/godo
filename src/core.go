package godo

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var addCommand *flag.FlagSet
var listCommand *flag.FlagSet
var completeCommand *flag.FlagSet
var removeCommand *flag.FlagSet

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

type completeCommandFlags struct {
	id int
}

type removeCommandFlags struct {
	id int
}

var addFlags addCommandFlags
var listFlags listCommandFlags
var completeFlags completeCommandFlags
var removeFlags removeCommandFlags

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

	completeCommand = flag.NewFlagSet("comp", flag.ExitOnError)

	completeFlags = completeCommandFlags{}

	removeCommand = flag.NewFlagSet("rm", flag.ExitOnError)

	removeFlags = removeCommandFlags{}

	addCommand.Usage = func() {
		ExitWithUsage()
	}

	listCommand.Usage = func() {
		ExitWithUsage()
	}

	completeCommand.Usage = func() {
		ExitWithUsage()
	}

	removeCommand.Usage = func() {
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
	case "comp":
		completeCommand.Parse(os.Args[2:])

		argument := completeCommand.Arg(0)

		if argument == "" {
			FatalWithUsage("Missing todo id")
			return
		}

		todoId, err := strconv.Atoi(completeCommand.Arg(0))

		if err != nil {
			log.Fatalln(err)
		}

		completeFlags.id = todoId

		complete(completeFlags)
		break
	case "rm":
		removeCommand.Parse(os.Args[2:])

		argument := removeCommand.Arg(0)

		if argument == "" {
			FatalWithUsage("Missing todo id")
			return
		}

		todoId, err := strconv.Atoi(removeCommand.Arg(0))

		if err != nil {
			log.Fatalln(err)
		}

		removeFlags.id = todoId

		remove(removeFlags)
		break
	}
}
