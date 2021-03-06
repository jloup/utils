package utils_test

import (
	"fmt"

	"github.com/jloup/utils"
)

func ExampleErrorAggregator() {

	var UserError = utils.New("UserError", 0)
	var ParsingError = utils.New("ParsingError", 1)
	var IOError = utils.New("IOError", 2)

	report := utils.NewErrorAggregator()

	report.New(UserError, "bad input")
	report.New(UserError, "not identified")
	report.New(ParsingError, "malformed")
	report.New(IOError, "cannot find resource")

	fmt.Printf("Report (all errors):\n%s\n\n", report.Error())

	fmt.Printf("Report (user errors):\n%s\n\n", report.ErrorWithCode(UserError).Error())

	fmt.Printf("Report (io and parsing errors):\n%s\n\n", report.ErrorWithCode(utils.Join("", ParsingError, IOError)).Error())

	//Output:
	//Report (all errors):
	//[UserError]
	//	bad input
	//	not identified
	//[ParsingError]
	//	malformed
	//[IOError]
	//	cannot find resource
	//
	//Report (user errors):
	//[UserError]
	//	bad input
	//	not identified
	//
	//Report (io and parsing errors):
	//[ParsingError]
	//	malformed
	//[IOError]
	//	cannot find resource
}
