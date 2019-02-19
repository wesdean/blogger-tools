package blogger_tools_lib

import (
	"errors"
	"github.com/google/logger"
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"io"
	"io/ioutil"
	"os"
)

type NotifyTool struct {
	Logger *logger.Logger
}

type NotifyToolArgs struct {
	Config      *Config
	Verbose     bool
	SysLog      bool
	LogFilePath string
}

func NewNotifyTool() *NotifyTool {
	return &NotifyTool{}
}

func (tool *NotifyTool) Run(args *NotifyToolArgs) (err error) {
	var logFile io.Writer
	if args.LogFilePath != "" {
		logFile, err = os.OpenFile(args.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	} else {
		logFile = ioutil.Discard
	}
	tool.Logger = logger.Init("NotifyToolLog", args.Verbose, args.SysLog, logFile)
	defer tool.Logger.Close()

	tool.Logger.Info("Running NotifyTool")

	var errorFlag = false
	for _, blogConfig := range args.Config.Blogs {
		tool.Logger.Infof("Running NotifyTool for BlogID: %s", blogConfig.ID)
		blogger := blogger_tools_blogger.NewBlogger(tool.Logger, blogConfig.AccessToken, blogConfig.ID)
		blog, err := blogger.Blog.Get()
		if err != nil {
			errorFlag = true
			tool.Logger.Errorf("BlogID: %s; Message: %s", blogConfig.ID, err)
			continue
		}

		tool.Logger.Info(blog)
	}
	if errorFlag {
		return errors.New("There were errors running NotifyTool")
	}

	tool.Logger.Info("NotifyTool done")

	return err
}
