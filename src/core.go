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
var overviewCommand *flag.FlagSet

type addCommandFlags struct {
	title       *string
	description *string
	priority    *int
	project     string
}

type listCommandFlags struct {
	showProjects *bool
	showAll      *bool
	project      string
}

type completeCommandFlags struct {
	id int64
}

type removeCommandFlags struct {
	id int64
}

type overviewCommandFlags struct {
	showAll *bool
}

var addFlags addCommandFlags
var listFlags listCommandFlags
var completeFlags completeCommandFlags
var removeFlags removeCommandFlags
var overviewFlags overviewCommandFlags

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

	listFlags.showAll = listCommand.Bool("all", false, "")
	listCommand.BoolVar(listFlags.showAll, "a", false, "")

	completeCommand = flag.NewFlagSet("comp", flag.ExitOnError)

	completeFlags = completeCommandFlags{}

	removeCommand = flag.NewFlagSet("rm", flag.ExitOnError)

	removeFlags = removeCommandFlags{}

	overviewCommand = flag.NewFlagSet("overview", flag.ExitOnError)

	overviewFlags = overviewCommandFlags{}
	overviewFlags.showAll = overviewCommand.Bool("all", false, "")
	overviewCommand.BoolVar(overviewFlags.showAll, "a", false, "")

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
	cleanup()

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

		completeFlags.id = int64(todoId)

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

		removeFlags.id = int64(todoId)

		remove(removeFlags)
		break
	case "overview", "ov":
		overviewCommand.Parse(os.Args[2:])

		overview(overviewFlags)
		break
	default:
		FatalWithUsage("Invalid subcommand")
		break
	}
}
