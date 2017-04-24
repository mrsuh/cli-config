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
}

var configInstasnce *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		configInstasnce = &singleton{}
	})
	return configInstasnce
}

func (s *singleton) Init() error {

	if len(s.data) > 0 {
		return nil
	}

	args := os.Args
	if len(args) < 2 {
		return errors.New("Config file is required")
	}

	data, read_err := ioutil.ReadFile(args[1])

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