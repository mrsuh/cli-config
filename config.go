package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"errors"
	"sync"
)

type singleton struct {
	data map[string]interface{}
	path string
}

var configInstance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		configInstance = &singleton{}
	})
	return configInstance
}

func (s *singleton) SetPath(path string) {
	s.path = path
}

func (s *singleton) Init() error {

	if len(s.data) > 0 {
		return nil
	}

	var path string
	if s.path != "" {
		path = s.path
	} else {
		args := os.Args
		if len(args) < 2 {
			return errors.New("Config file is required")
		}
		path = args[1]
	}

	if path == "" {
		return errors.New("Config file is required")
	}

	data, read_err := ioutil.ReadFile(path)

	if read_err != nil {
		return read_err
	}

	parse_err := yaml.Unmarshal(data, &s.data)
	if parse_err != nil {
		return parse_err
	}

	return nil
}

func (s singleton) Get() map[string]interface{} {
	return s.data
}