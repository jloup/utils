package utils

type ErrorFlagged interface {
	error
	Flag() Flag
	FlagString() string
	Msg() string
	ErrorWithCode(Flag) ErrorFlagged
}
