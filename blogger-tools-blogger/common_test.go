package blogger_tools_blogger_test

import (
	"github.com/google/logger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"io/ioutil"
)

var skipExternalServices = true

func getConfig() (*blogger_tools_lib.Config, error) {
	accessToken, err := ioutil.ReadFile("../secrets/access_token.txt")
	if err != nil {
		return nil, err
	}

	return &blogger_tools_lib.Config{
		Environment: "test",
		Blogs: []blogger_tools_lib.BlogConfig{
			{
				AccessToken: string(accessToken),
				ID:          "3051261493420306591",
			},
		},
	}, nil
}

func getLogger() *logger.Logger {
	return logger.Init("Test Logger", true, false, ioutil.Discard)
}
