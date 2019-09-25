# expecting

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
parser.Description("some description")

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

// dump unexpected ars into var
parser.Extras(&extras)

if err := parser.Parse(os.Args); nil != err {
    fmt.Println(parser.Help())
}

```
