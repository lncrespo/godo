package godo

import (
	"flag"
	"os"

	"github.com/lncrespo/godo/src/migrations"
)

var addCommand *flag.FlagSet
var listCommand *flag.FlagSet
var completeCommand *flag.FlagSet
var removeCommand *flag.FlagSet
var overviewCommand *flag.FlagSet
var infoCommand *flag.FlagSet

var addFlags addCommandFlags
var listFlags listCommandFlags
var completeFlags completeCommandFlags
var removeFlags removeCommandFlags
var overviewFlags overviewCommandFlags
var infoFlags infoCommandFlags

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

	addFlags.dueAt = addCommand.String("due-at", "", "")
	addCommand.StringVar(addFlags.dueAt, "D", "", "")

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

	infoCommand = flag.NewFlagSet("info", flag.ExitOnError)

	infoFlags = infoCommandFlags{}

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
	migrations.ExecuteMigrations()
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

		todoId, err := getIdFromArgument(completeCommand.Arg(0))

		if err != nil {
			FatalWithUsage(err.Error())
		}

		completeFlags.id = todoId

		complete(completeFlags)
		break
	case "rm":
		removeCommand.Parse(os.Args[2:])

		todoId, err := getIdFromArgument(removeCommand.Arg(0))

		if err != nil {
			FatalWithUsage(err.Error())
		}

		removeFlags.id = todoId

		remove(removeFlags)
		break
	case "overview", "ov":
		overviewCommand.Parse(os.Args[2:])

		overview(overviewFlags)
		break
	case "info":
		infoCommand.Parse(os.Args[2:])

		todoId, err := getIdFromArgument(infoCommand.Arg(0))

		if err != nil {
			FatalWithUsage(err.Error())
		}

		infoFlags.id = todoId

		info(infoFlags)

		break
	case "reset":
		reset()

		break
	default:
		FatalWithUsage("Invalid subcommand")
		break
	}
}
