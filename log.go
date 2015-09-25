package utils

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Error struct {
	Flag    Flag
	Message string
}

func (e Error) FlagString() string {
	return e.Flag.String()
}

func (e Error) Msg() string {
	return e.Message
}

func (e Error) Error() string {
	return e.Message
}

func ErrIs(err error, f Flag) bool {
	switch e := err.(type) {
	case Error:
		return Intersect(e.Flag, f)
	case *Error:
		return Intersect(e.Flag, f)
	}

	return false
}

type L struct {
	*log.Entry
	C Counter
}

const (
	fieldFlagName      = "flag"
	fieldFileInnfoName = "loc"
)

func fileinfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func bindType(fields log.Fields, args ...interface{}) {
	for _, arg := range args {
		switch e := arg.(type) {
		case *Error:
			fields[fieldFlagName] = e.Flag.String()
		case Error:
			fields[fieldFlagName] = e.Flag.String()
		}
	}
}

func (l *L) NewErrorFlag(name string) Flag {
	return InitFlag(&l.C, name)
}

func (l *L) Error(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Error(args...)
}

func (l *L) Errorf(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Errorf(f, args...)
}

func (l *L) Errorln(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Errorln(args...)
}

func (l *L) Debug(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Debug(args...)
}

func (l *L) Debugf(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Debugf(f, args...)
}

func (l *L) Debugln(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Debugln(args...)
}

func (l *L) Fatal(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Fatal(args...)
}

func (l *L) Fatalf(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Fatalf(f, args...)
}

func (l *L) Fatalln(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Fatalln(args...)
}

func (l *L) Panic(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Panic(args...)
}

func (l *L) Panicf(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Panicf(f, args...)
}

func (l *L) Panicln(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Panicln(args...)
}

func (l *L) Print(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Print(args...)
}

func (l *L) Printf(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Printf(f, args...)
}

func (l *L) Println(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Println(args...)
}

func (l *L) Warn(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Warn(args...)
}

func (l *L) Warnf(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Warnf(f, args...)
}

func (l *L) Warnln(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Warnln(args...)
}

func (l *L) Info(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Info(args...)
}

func (l *L) Infof(f string, args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Infof(f, args...)
}

func (l *L) Infoln(args ...interface{}) {
	fields := log.Fields{fieldFileInnfoName: fileinfo(2)}

	bindType(fields, args...)

	l.Entry.WithFields(fields).Infoln(args...)
}

func (l *L) WithField(key string, value interface{}) *L {
	return &L{l.Entry.WithField(key, value), l.C}
}

func (l *L) WithFields(fields log.Fields) *L {
	return &L{l.Entry.WithFields(fields), l.C}
}

func Concat(flags ...Flag) Flag {
	var s []string
	for _, f := range flags {
		s = append(s, f.String())
	}

	return Join(strings.Join(s, "|"), flags...)
}
