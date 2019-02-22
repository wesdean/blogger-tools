package blogger_tools_lib

import (
	"encoding/json"
	"github.com/google/logger"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Environment      string
	LogDirectory     string
	SecretsDirectory string
	Blogger          *BloggerConfig
	SendGrid         *SendGridConfig
	Logs             *BlogConfigLogs
	NotifyTool       *NotifyToolConfig
}

type BloggerConfig struct {
	Blogs []BlogConfig
}

type BlogConfig struct {
	AccessTokenFile *string
	AccessToken     *string
	ID              string
}

type BlogConfigLogs struct {
	General    string
	NotifyTool string
}

type SendGridConfig struct {
	APIKey           string
	DefaultFromEmail string
	DefaultFromName  string
}

type NotifyToolConfig struct {
	BlogUpdatedRecipientsFile string
}

func NewConfig(fileName string) (*Config, error) {
	var config = &Config{
		LogDirectory:     "./logs",
		SecretsDirectory: "./secrets",
		Logs: &BlogConfigLogs{
			NotifyTool: "notify-tool.log",
		},
	}

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

	logDirectory, err := filepath.Abs(config.LogDirectory)
	if err != nil {
		return nil, err
	}
	config.LogDirectory = logDirectory
	if _, err := os.Stat(config.LogDirectory); os.IsNotExist(err) {
		err = os.Mkdir(config.LogDirectory, 0666)
	}
	if err != nil {
		return nil, err
	}

	secretsDirectory, err := filepath.Abs(config.SecretsDirectory)
	if err != nil {
		return nil, err
	}
	config.SecretsDirectory = secretsDirectory
	if _, err := os.Stat(config.SecretsDirectory); os.IsNotExist(err) {
		err = os.Mkdir(config.SecretsDirectory, 0666)
	}
	if err != nil {
		return nil, err
	}

	for index := range config.Blogger.Blogs {
		blogConfig := &config.Blogger.Blogs[index]

		if blogConfig.AccessToken == nil {
			accessToken, err := ioutil.ReadFile(config.BuildSecretFilePath(*blogConfig.AccessTokenFile))
			if err != nil {
				return nil, err
			}
			tokenString := string(accessToken)
			blogConfig.AccessToken = &tokenString
		}
	}

	return config, nil
}

func (config *Config) BuildLogFilePath(filename string) string {
	filename = config.LogDirectory + "/" + filename
	dirName := filepath.Dir(filename)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0666)
	}
	return filepath.FromSlash(filename)
}

func (config *Config) BuildSecretFilePath(filename string) string {
	filename = config.SecretsDirectory + "/" + filename
	dirName := filepath.Dir(filename)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0666)
	}
	return filepath.FromSlash(filename)
}

func (config *Config) CreateLogger(filename string, reset bool) (*logger.Logger, error) {
	var logFile io.Writer
	var err error

	if filename != "" {
		writeType := os.O_APPEND
		if reset {
			writeType = os.O_TRUNC
		}
		logFile, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|writeType, 0600)
		if err != nil {
			return nil, err
		}
	} else {
		logFile = ioutil.Discard
	}
	return logger.Init("NotifyToolLog", false, false, logFile), nil
}
