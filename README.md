# godo

A simple todo-list program written in Go.

## Requirements

- Sqlite

This program uses sqlite(3) to store todos. Make sure sqlite is installed.

## Usage

### Add a todo

```
godo add [flags] <project>
```

#### Optional flags

|Flag|Description|
|----|-----------|
|`-t, --title`|Specify a title for the todo. If this parameter is not given, the program will launch interactive mode.|
|`-d, --description`|Specify a description for the todo|
|`-p, --priority`|Specify the priority for the todo (0-9)|
|`-D, --due-at`|Set a due date for the todo. Use a time formatted string in the format "d-m-Y hh:ss". Examples: "11-4-2022 13:00", "02-04-2023 12:20|

#### Optional arguments

|Argument|Description|
|--------|-----------|
|`project`|Add the todo in the given project. If it doesn't exist, the project will be automatically created|

### List todos

```
godo list [flags] <project>
```

#### Optional flags

|Flag|Description|
|----|-----------|
|`-a, --all`|Include already completed todos|
|`-p, --projects`|List all projects|

#### Optional arguments

|Argument|Description|
|--------|-----------|
|`project`|List todos for the given project|

### Complete a todo

```
godo comp <id>
```

#### Required arguments

|Argument|Description|
|--------|-----------|
|`id`|The id of the todo to be completed. Ids are shown when using `godo list` or `godo ov(erview)`|

### Remove a todo

```
godo rm <id>
```

#### Required arguments

|Argument|Description|
|--------|-----------|
|`id`|The id of the todo to be removed. Ids are shown when using `godo list` or `godo ov(erview)`|

### Show an overview

```
godo ov
```

```
godo overview
```

#### Optional flags

|Flag|Description|
|----|-----------|
|`-a, --all`|Include already completed todos to overview|

### Show details of a todo

```
godo info <id>
```

#### Required arguments

|Argument|Description|
|--------|-----------|
|`id`|The id of the todo to be shown. Ids are shown when using `godo list` or `godo ov(erview)`|

### Remove everything

```
godo reset
```

Resets the state of the database, removing every project and todo. This action cannot be reversed.
When executing, you will be prompted for confirmation.
