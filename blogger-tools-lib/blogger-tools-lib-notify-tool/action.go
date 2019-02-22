package blogger_tools_lib_notify_tool

import (
	"github.com/google/logger"
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
)

type Action struct {
	Config          *blogger_tools_lib.Config
	Logger          *logger.Logger
	Blog            *blogger_tools_blogger.Blog
	BlogAccessToken string
}
