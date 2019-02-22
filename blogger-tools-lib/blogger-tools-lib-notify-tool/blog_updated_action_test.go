package blogger_tools_lib_notify_tool_test

import (
	"encoding/json"
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"github.com/wesdean/blogger-tools/blogger-tools-lib/blogger-tools-lib-notify-tool"
	"io/ioutil"
	"testing"
)

func TestBlogUpdatedAction_Do(t *testing.T) {
	config, err := blogger_tools_lib.NewConfig(notifyToolConfigFile)
	if err != nil {
		t.Error(err)
		return
	}

	log, err := config.CreateLogger("../../logs/notify-tool-test.log", true)
	if err != nil {
		t.Error(err)
		return
	}

	blogger := blogger_tools_blogger.NewBlogger(log, *config.Blogger.Blogs[0].AccessToken, config.Blogger.Blogs[0].ID)
	blog, err := blogger.Blog.Get()
	if err != nil {
		t.Error(err)
		return
	}

	recipientFile := config.BuildSecretFilePath(config.NotifyTool.BlogUpdatedRecipientsFile)
	recipientJSON, err := ioutil.ReadFile(recipientFile)
	if err != nil {
		t.Error(err)
		return
	}

	var allRecipients map[string][]blogger_tools_lib_notify_tool.BlogUpdatedRecipient
	err = json.Unmarshal(recipientJSON, &allRecipients)
	if err != nil {
		t.Error(err)
		return
	}

	notifyAction := &blogger_tools_lib_notify_tool.Action{
		Config:          config,
		Logger:          log,
		Blog:            blog,
		BlogAccessToken: *config.Blogger.Blogs[0].AccessToken,
	}
	action := &blogger_tools_lib_notify_tool.BlogUpdatedAction{
		notifyAction,
		allRecipients[blog.Id],
	}
	err = action.Do()
	if err != nil {
		t.Error(err)
		return
	}
}
