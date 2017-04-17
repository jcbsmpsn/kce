package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var Config Configuration

type Configuration struct {
	Core    Core
	Include Include
	Default Default
	Alias   map[string]string
}

type Core struct {
	Kubectl string
}

type Include struct {
	Include string
}

type Default struct {
	Namespace string
}

// Load loads up configuration from the config file in the users home
// directory, and writes a default config file there if there currently is no
// config file.
func Load() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	configFile := filepath.Join(usr.HomeDir, ".kce.config")
	if _, err := os.Stat(configFile); err == nil {
		if _, err := toml.DecodeFile(configFile, &Config); err != nil {
			return fmt.Errorf("Could not load the config file %q: %v\n", configFile, err)
		}
	} else if os.IsNotExist(err) {
		if err := ioutil.WriteFile(configFile, []byte(DefaultConfigFileText), 0600); err != nil {
			return fmt.Errorf("Unable to write a default config file to %q: %v", configFile, err)
		}
	} else {
		return fmt.Errorf("Unable to read config file %q: %v", configFile, err)
	}
	return nil
}
