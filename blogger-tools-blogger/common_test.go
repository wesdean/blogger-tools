package blogger_tools_blogger_test

import "github.com/wesdean/blogger-tools/blogger-tools-lib"

func getConfig() *blogger_tools_lib.Config {
	return &blogger_tools_lib.Config{
		Environment: "test",
		APIKey:      "AIzaSyC6QvuYH5rA-NExrFnStRisKARjmcar3QU",
		BlogIDs:     []string{"3960547499512363533"},
	}
}
