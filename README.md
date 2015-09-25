# Flag

Simple flag utility 

Can be used for error customizing, debug level handling

Example:
```go
import (
    "fmt"
    "github.com/jloup/flag"
)

// we define our flags
var (
    BadInput      = flag.New("BadInput", 1)
    NotAuthorized = flag.New("NotAuthorized", 2)
    InternalError = flag.New("InternalError", 3)
    // aggregat of flags
    UserError = flag.Join("UserError", BadInput, NotAuthorized)
)

// error flagged
type MyError struct {
    flag flag.Flag
    Msg  string
}
// MyError with flag BadInput
match := MyError{flag: BadInput, Msg: "json wanted"}
// MyError with flag InternalError
dontmatch := MyError{flag: InternalError, Msg: "server error"}

// will print
if flag.Intersect(match.flag, UserError) {
    fmt.Printf("user error happened '[%s] %s'\n",
        match.flag.String(), match.Msg)
}

// will not print
if flag.Intersect(dontmatch.flag, UserError) {
    fmt.Printf("user error happened '[%s] %s'\n",
        dontmatch.flag.String(), dontmatch.Msg)
}
```
Output:
```
user error happened '[BadInput] json wanted'
```

## flag.Counter
Instead of flag.New, flag.InitFlag can be employed to avoid the pain to manage numerical values
```go
import (
    "fmt"
    "github.com/jloup/flag"
)

// we define our flags
var (
    c flag.Counter = 0
    BadInput       = flag.InitFlag(&c, "BadInput")
    NotAuthorized  = flag.InitFlag(&c, "NotAuthorized")
    InternalError  = flag.InitFlag(&c, "InternalError")
    // aggregat of flags
    UserError = flag.Join("UserError", BadInput, NotAuthorized)
)

// error flagged
type MyError struct {
    flag flag.Flag
    Msg  string
}
// MyError with flag BadInput
match := MyError{flag: BadInput, Msg: "json wanted"}
// MyError with flag InternalError
dontmatch := MyError{flag: InternalError, Msg: "server error"}

// will print
if flag.Intersect(match.flag, UserError) {
    fmt.Printf("user error happened '[%s] %s'\n",
        match.flag.String(), match.Msg)
}

// will not print
if flag.Intersect(dontmatch.flag, UserError) {
    fmt.Printf("user error happened '[%s] %s'\n",
        dontmatch.flag.String(), dontmatch.Msg)
}
```
Output:
```
user error happened '[BadInput] json wanted'
```
## errors

Package errors provides functions to easily aggregate, filter and pretty print errors with different type

Example:
```go
import (
	"fmt"
	"github.com/jloup/errors"
	"github.com/jloup/flag"
)

func main() {
var UserError = flag.New("UserError", 0)
var ParsingError = flag.New("ParsingError", 1)
var IOError = flag.New("IOError", 2)

report := errors.NewErrorAggregator()

report.New(UserError, "bad input")
report.New(UserError, "not identified")
report.New(ParsingError, "malformed")
report.New(IOError, "cannot find resource")

fmt.Printf("Report (all errors):\n%s\n", 
report.Error())

fmt.Printf("Report (user errors):\n%s\n\n", 
report.ErrorWithCode(UserError).Error())

fmt.Printf("Report (io and parsing errors):\n%s\n\n", 
report.ErrorWithCode(flag.Join("", ParsingError, IOError)).Error())
}
```
Output:
```
Report (all errors):
[UserError]
	bad input
	not identified
[ParsingError]
	malformed
[IOError]
	cannot find resource

Report (user errors):
[UserError]
	bad input
	not identified

Report (io and parsing errors):
[ParsingError]
	malformed
[IOError]
	cannot find resource
```
