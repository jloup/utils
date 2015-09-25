package utils_test

import (
	"fmt"

	"github.com/jloup/utils"
)

func ExampleFlag() {
	// we define our flags
	var (
		BadInput      = utils.New("BadInput", 1)
		NotAuthorized = utils.New("NotAuthorized", 2)
		InternalError = utils.New("InternalError", 3)
		// aggregat of flags
		UserError = utils.Join("UserError", BadInput, NotAuthorized)
	)

	// error flagged
	type MyError struct {
		flag utils.Flag
		Msg  string
	}
	// MyError with flag BadInput
	match := MyError{flag: BadInput, Msg: "json wanted"}
	// MyError with flag InternalError
	dontmatch := MyError{flag: InternalError, Msg: "server error"}

	// will print
	if utils.Intersect(match.flag, UserError) {
		fmt.Printf("user error happened '[%s] %s'\n",
			match.flag.String(), match.Msg)
	}

	// will not print
	if utils.Intersect(dontmatch.flag, UserError) {
		fmt.Printf("user error happened '[%s] %s'\n",
			dontmatch.flag.String(), dontmatch.Msg)
	}
	// Output:
	// user error happened '[BadInput] json wanted'
}

func ExampleCounter() {
	var c utils.Counter = 0

	var (
		BadInput      = utils.InitFlag(&c, "BadInput")
		NotAuthorized = utils.InitFlag(&c, "NotAuthorized")
		InternalError = utils.InitFlag(&c, "InternalError")
	)

	if !utils.Intersect(BadInput, NotAuthorized) && !utils.Intersect(NotAuthorized, InternalError) && !utils.Intersect(InternalError, BadInput) {
		fmt.Printf("None of them match!")
	}

	// Output:
	// None of them match!
}
