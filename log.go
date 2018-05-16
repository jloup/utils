package utils

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Error struct {
	ErrorFlag Flag
	Message   string
}

func (e Error) FlagString() string {
	return e.ErrorFlag.String()
}

func (e Error) Flag() Flag {
	return e.ErrorFlag
}

func (e Error) Msg() string {
	return e.Message
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) ErrorWithCode(f Flag) ErrorFlagged {
	if Intersect(e.ErrorFlag, f) {
		return e
	}

	return nil
}

func ErrIs(err error, f Flag) bool {
	if err == nil {
		return false
	}

	switch e := err.(type) {
	case ErrorFlagged:
		if Intersect(e.Flag(), f) {
			return f.name == e.FlagString()
		}
	}

	return false
}

var ErrorFlagCounter Counter

func NewErrorFlag(name string) Flag {
	return InitFlag(&ErrorFlagCounter, name)
}

type L struct {
	*log.Entry
	configuration LogConfiguration
}

var stdL *L

func init() {
	stdL = &L{Entry: log.NewEntry(log.StandardLogger())}
	stdL.configuration.SetToDefaults()
}

func StandardL() *L {
	return stdL
}

type LogConfiguration struct {
	LogOut      string `toml:"log_out"`
	LogLevel    string `toml:"log_level"`
	LogColors   bool   `toml:"log_colors"`
}

func (l *LogConfiguration) SetToDefaults() {
	*l = LogConfiguration{"stdout", "error", false}
}

func ConfigureStdLogger(lc LogConfiguration) error {
	var l log.Level
	var o io.Writer
	var err error

	switch lc.LogOut {
	case "stdout":
		o = os.Stdout
	default:
		o, err = os.OpenFile(lc.LogOut, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	}

	if err != nil {
		return err
	}

	switch lc.LogLevel {
	case "debug":
		l = log.DebugLevel
	case "info":
		l = log.InfoLevel
	case "warn":
		l = log.WarnLevel
	case "error":
		l = log.ErrorLevel
	case "fatal":
		l = log.FatalLevel
	case "panic":
		l = log.PanicLevel
	default:
		err = fmt.Errorf("log level not recoginzed '%v'", lc.LogLevel)
	}

	if err != nil {
		return err
	}

	stdL.Logger.Out = o
	stdL.Logger.Formatter = &log.TextFormatter{
		DisableColors:    !lc.LogColors,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "Jan _2 15:04:05"}

	stdL.Logger.Level = l
	stdL.configuration = lc

	return nil
}

const (
	fieldFlagName      = "flag"
	fieldFileInnfoName = "loc"
)

func bindType(fields log.Fields, args ...interface{}) {
	for _, arg := range args {
		switch e := arg.(type) {
		case ErrorFlagged:
			fields[fieldFlagName] = e.FlagString()
		}
	}
}

func (l *L) Error(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Error(args...)
}

func (l *L) Errorf(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Errorf(f, args...)
}

func (l *L) Errorln(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Errorln(args...)
}

func (l *L) Debug(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Debug(args...)
}

func (l *L) Debugf(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Debugf(f, args...)
}

func (l *L) Debugln(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Debugln(args...)
}

func (l *L) Fatal(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Fatal(args...)
}

func (l *L) Fatalf(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Fatalf(f, args...)
}

func (l *L) Fatalln(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Fatalln(args...)
}

func (l *L) Panic(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Panic(args...)
}

func (l *L) Panicf(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Panicf(f, args...)
}

func (l *L) Panicln(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Panicln(args...)
}

func (l *L) Print(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Print(args...)
}

func (l *L) Printf(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Printf(f, args...)
}

func (l *L) Println(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Println(args...)
}

func (l *L) Warn(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Warn(args...)
}

func (l *L) Warnf(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Warnf(f, args...)
}

func (l *L) Warnln(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Warnln(args...)
}

func (l *L) Info(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Info(args...)
}

func (l *L) Infof(f string, args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Infof(f, args...)
}

func (l *L) Infoln(args ...interface{}) {
	fields := log.Fields{}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Infoln(args...)
}

func (l *L) WithField(key string, value interface{}) *L {
	return &L{Entry: l.Entry.WithField(key, value),
		configuration: l.configuration,
	}
}

func (l *L) WithFields(fields log.Fields) *L {
	return &L{Entry: l.Entry.WithFields(fields),
		configuration: l.configuration,
	}
}

func Concat(flags ...Flag) Flag {
	var s []string
	for _, f := range flags {
		s = append(s, f.String())
	}

	return Join(strings.Join(s, "|"), flags...)
}
