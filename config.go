package utils

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
)

type Conf interface {
	SetToDefaults()
}

func MustParseStdConfigFile(conf Conf) {
	var path string

	flag.StringVar(&path, "conf", "conf.toml", "config file")
	flag.Parse()

	MustParseConfigFile(path, conf)
}

func MustParseConfigFile(path string, conf Conf) {
	conf.SetToDefaults()
	MustDecodeConfigFile(path, conf)
}

func MustDecodeConfigFile(path string, conf interface{}) {
	m, err := toml.DecodeFile(path, conf)
	if err != nil {
		panic(err)
	}

	s := reflect.TypeOf(conf).Elem()
	var missingParams []string
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		if field.Tag.Get("conf") != "required" {
			continue
		}
		if !m.IsDefined(strings.ToLower(field.Name)) {
			missingParams = append(missingParams, strings.ToLower(field.Name))
		}
	}

	if len(missingParams) != 0 {
		panic(fmt.Sprintf("following config parameters are missing: %s", strings.Join(missingParams, ", ")))
	}
}
