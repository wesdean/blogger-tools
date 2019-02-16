package blogger_tools_lib

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Environment string
	APIKey      string
	BlogIDs     []string
}

func NewConfig(fileName string) (*Config, error) {
	var config = &Config{}

	if fileName != "" {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
