# arg
argument parsing lib

Note: this is a work in progress and is not yet operational

## TODO
- [ ] need a better way of handling =\
  unexpected args dont respect it for dashed and groups
- [x] remove preceeding - and -- from unexpected args
- [x] group parsing is basically fucked
- [x] commands
- [.] generate help

## Api

```go
import (
    "fmt"
    "arg"
    "os"
)

var str string
var strs []string
var extras []arg.Argument

parser = arg.NewParser()

parser.Version("0.1.0")
parser.Title("Thing")
parser.Usage("some description")

parser.String(&str, "s,string")
parser.Strings(&srts, "strings,S,all-the-strings", &arg.Meta{
    Usage: "the thing",
    Default: []string{"thing"},
    Required: false,
    Override: false,
    Validate: func(arg string) error {
        ...
    }
}

// commands do not actually get validated or parsed in any way
// they exist purely for the creating help documentation
parser.Command("command-name", &arg.CommandConfig{
    Usage: "this is a sub command",
    Title: "Sub Command",
})

// dump unexpected ars into var
parser.Extras(&extras)

if err := parser.Parse(os.Args); nil != err {
    fmt.Println(parser.Help())
}

```
