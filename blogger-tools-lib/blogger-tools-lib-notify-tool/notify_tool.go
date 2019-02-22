package blogger_tools_lib_notify_tool

import (
	"errors"
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
)

type NotifyTool struct {
	*blogger_tools_lib.BloggerTool
}

type NotifyToolArgs struct {
	ResetLog bool
	Actions  *NotifyToolActions
}

type NotifyToolActions struct {
	SendLatestPost *SendLatestPostOptions
}

type SendLatestPostOptions struct {
}

func (tool *NotifyTool) Run(args *NotifyToolArgs) (err error) {
	log, err := tool.Config.CreateLogger(tool.Config.BuildLogFilePath(tool.Config.Logs.NotifyTool), args.ResetLog)
	if err != nil {
		return err
	}
	defer log.Close()

	log.Info("Running NotifyTool")

	var errorFlag = false
	for _, blogConfig := range tool.Config.Blogger.Blogs {
		log.Infof("Running NotifyTool for BlogID: %s", blogConfig.ID)
		if blogConfig.AccessToken == nil {
			errorFlag = true
			log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, "Missing access token")
		}
		blogger := blogger_tools_blogger.NewBlogger(log, *blogConfig.AccessToken, blogConfig.ID)
		blog, err := blogger.Blog.Get()
		if err != nil {
			errorFlag = true
			log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, err)
			continue
		}

		if args.Actions == nil {
			log.Infof("Blog check successful for %v", blog.Id)
		} else {

		}
	}
	if errorFlag {
		return errors.New("There were errors running NotifyTool")
	}

	log.Info("NotifyTool done")

	return err
}
