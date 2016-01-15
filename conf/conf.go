package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration is the struct representing a configuration
type Configuration struct {
	Port       int  `yaml:"port"`
	Debug      bool `yaml:"debug"`
	APIVersion int  `yaml:"api_version"`
}

// C is the main Configuration
var C Configuration

// Load loads the given fp (file path) to the C global configuration variable.
func Load(fp string) error {
	var err error
	conf, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(conf, &C)
	return err
}
