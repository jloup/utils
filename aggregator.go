package utils

import "fmt"

type element struct {
	flag  Flag
	error string
}

// ErrorAggregator allows to gather errors from different sources in one place
type ErrorAggregator struct {
	flags  Flag
	errors []element
}

func NewErrorAggregator() ErrorAggregator {
	return ErrorAggregator{flags: Flag{}, errors: make([]element, 0)}
}

// Add a new error to aggregat providing its flag and msg
func (e *ErrorAggregator) New(flags Flag, msg string) {

	var index int
	if index = e.findIndex(flags); index == -1 {
		index = e.createElement(flags)
		s := e.flags.String()
		if s != "" {
			s += "/"
		}
		s += flags.String()
		e.flags = Join(s, e.flags, flags)
	} else {
		e.flags = Join(e.FlagString(), e.flags, flags)
	}

	if e.errors[index].error != "" {
		e.errors[index].error += "\n\t"
	}
	e.errors[index].error += fmt.Sprintf("%s", msg)
}

func (e *ErrorAggregator) findIndex(f Flag) int {
	for i, el := range e.errors {
		if f.N.Cmp(&el.flag.N) == 0 {
			return i
		}
	}
	return -1
}

func (e *ErrorAggregator) createElement(f Flag) int {
	if index := e.findIndex(f); index == -1 {
		e.errors = append(e.errors, element{flag: f})
		return len(e.errors) - 1

	} else {
		return index
	}
}

func (e *ErrorAggregator) FlagString() string {
	return e.flags.String()
}

// Add a new error to aggregat given an ErrorFlagged itf
func (e *ErrorAggregator) NewError(err ErrorFlagged) {
	e.New(err.Flag(), err.Msg())
}

func (e *ErrorAggregator) Flag() Flag {
	return e.flags
}

// returns nil if aggregat if empty
func (e *ErrorAggregator) ErrorObject() ErrorFlagged {
	if e.flags.N.Cmp(&zero) != 0 {
		return e
	}
	return nil
}

func (e *ErrorAggregator) ErrorWithCode(f Flag) ErrorFlagged {
	errAgg := NewErrorAggregator()
	if e.HasError(f) {
		for _, error := range e.errors {

			if Intersect(error.flag, f) {
				errAgg.New(error.flag, error.error)
			}
		}
		return &errAgg
	}
	return nil
}

func (e *ErrorAggregator) HasError(f Flag) bool {
	return Intersect(e.flags, f)
}

func (e *ErrorAggregator) filterMsg(f Flag) string {
	s := ""
	for _, error := range e.errors {

		if Intersect(error.flag, f) {
			if s != "" {
				s += "\n\t"
			}
			s += error.error
		}
	}
	return s
}

func (e *ErrorAggregator) Msg() string {
	s := ""
	for _, error := range e.errors {
		if s != "" {
			s += "\n\t"
		}
		s += error.error
	}
	return s
}

func (e *ErrorAggregator) Error() string {
	s := ""
	for _, error := range e.errors {

		if s != "" {
			s += "\n"
		}
		s += "[" + error.flag.String() + "]\n\t"
		s += error.error
	}
	return s

}
