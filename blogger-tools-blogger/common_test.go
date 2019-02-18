package blogger_tools_blogger_test

import (
	"github.com/google/logger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"io/ioutil"
)

func getConfig() *blogger_tools_lib.Config {
	return &blogger_tools_lib.Config{
		Environment: "test",
		Blogs: []blogger_tools_lib.BlogConfig{
			{
				APIKey: "ya29.GluzBhckAT2UMnmgrlYy8Usdbs22TNhNmwlZMFxUw7XZrOAnW0PA0S-G1QxA2jXSNrKQ__U6y4RTLkA8_r0mxnd11H8NsX1Xdu8Z4rj9gZkJu2lMsZ6rBxcQujNX",
				ID:     "1038304627327029055",
			},
		},
	}
}

func getLogger() *logger.Logger {
	return logger.Init("Test Logger", true, false, ioutil.Discard)
}
