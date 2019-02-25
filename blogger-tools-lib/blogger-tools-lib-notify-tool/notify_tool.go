package blogger_tools_lib_notify_tool

import (
	"encoding/json"
	"errors"
	"github.com/google/logger"
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"github.com/wesdean/blogger-tools/blogger-tools-lib/blogger-tools-lib-oauth-tool"
	"io/ioutil"
)

type NotifyTool struct {
	*blogger_tools_lib.BloggerTool
}

type NotifyToolArgs struct {
	ResetLog         bool
	Actions          *NotifyToolActions
	OAuthFlowHandler string
}

type NotifyToolResults struct {
	Success bool
	Actions NotifyToolActionResults
}

type NotifyToolActionResults struct {
	BlogUpdated bool
}

type NotifyToolActions struct {
	BlogUpdated *ActionBlogUpdatedOptions
}

type ActionBlogUpdatedOptions struct {
}

func (tool *NotifyTool) Run(args *NotifyToolArgs) (results *NotifyToolResults, err error) {
	results = &NotifyToolResults{}

	log, err := tool.Config.CreateLogger(tool.Config.BuildLogFilePath(tool.Config.Logs.NotifyTool), args.ResetLog)
	if err != nil {
		return results, err
	}
	defer log.Close()

	log.Info("Running NotifyTool")

	var blogUpdatedAllRecipients *map[string][]BlogUpdatedRecipient

	var errorFlag = false
	for index := range tool.Config.Blogger.Blogs {
		blogConfig := &tool.Config.Blogger.Blogs[index]

		log.Infof("Running NotifyTool for BlogID: %s", blogConfig.ID)
		if blogConfig.AccessToken == nil {
			err = tool.refreshAccessToken(log, blogConfig, args.OAuthFlowHandler)
			if err != nil {
				errorFlag = true
				log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, "Missing access token")
				continue
			}
		}
		blogger := blogger_tools_blogger.NewBlogger(log, *blogConfig.AccessToken, blogConfig.ID)
		blog, err, response := blogger.Blog.Get()
		if response != nil {
			if response.Error.Code == 401 || response.Error.Code == 403 {
				err = tool.refreshAccessToken(log, blogConfig, args.OAuthFlowHandler)
				if err == nil {
					blogger = blogger_tools_blogger.NewBlogger(log, *blogConfig.AccessToken, blogConfig.ID)
					blog, err, response = blogger.Blog.Get()
				}
			}
		}
		if err != nil {
			errorFlag = true
			log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, err)
			continue
		}

		if args.Actions == nil {
			log.Infof("Blog check successful for %v", blog.Id)
		} else {
			notifyAction := &Action{
				Config:          tool.Config,
				Logger:          log,
				Blog:            blog,
				BlogAccessToken: *tool.Config.Blogger.Blogs[0].AccessToken,
			}

			if args.Actions.BlogUpdated != nil {
				if blogUpdatedAllRecipients == nil {
					recipientFile := tool.Config.BuildSecretFilePath(tool.Config.NotifyTool.BlogUpdatedRecipientsFile)
					recipientJSON, err := ioutil.ReadFile(recipientFile)
					if err != nil {
						errorFlag = true
						log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, err)
						continue
					}

					err = json.Unmarshal(recipientJSON, &blogUpdatedAllRecipients)
					if err != nil {
						errorFlag = true
						log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, err)
						continue
					}
				}
			}
			action := &BlogUpdatedAction{
				notifyAction,
				(*blogUpdatedAllRecipients)[blog.Id],
			}
			err = action.Do()
			if err != nil {
				errorFlag = true
				log.Errorf("BlogID: %s; Message: %s", blogConfig.ID, err)
				continue
			}
			results.Actions.BlogUpdated = true
		}
	}
	if errorFlag {
		return results, errors.New("There were errors running NotifyTool")
	} else {
		results.Success = true
	}

	log.Info("NotifyTool done")

	return results, err
}

func (tool *NotifyTool) refreshAccessToken(log *logger.Logger, blogConfig *blogger_tools_lib.BlogConfig, flowHandlerName string) error {
	oauthTool := &blogger_tools_lib_oauth_tool.OAuthTool{tool.BloggerTool}
	_, err := oauthTool.RunForBlog(blogConfig, flowHandlerName)
	if err == nil {
		err = tool.Config.RefreshAccessToken(blogConfig)
	}
	return err
}
