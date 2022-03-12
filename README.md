# godo

A simple todo-list program written in Go.

## Usage

### Add a Todo  

```
godo add [flags] <project>
```

#### Optional flags

|Flag|Description|
|----|-----------|
|-t, --title|Specify a title for the todo. If this parameter is not given, the program will launch interactive mode.|
|-d, --description|Specify a description for the todo|
|-p, --priority|Specify the priority for the todo (0-9)|

#### Optional arguments

|Argument|Description|
|--------|-----------|
|project|Add the todo in the given project. If it doesn't exist, the project will be automatically created|

### List todos

```
godo list [flags] <project>
```

|Flag|Description|
|----|-----------|
|-p, --projects|List all projects|

#### Optional arguments

|Argument|Description|
|--------|-----------|
|project|List todos for the given project|
