package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"testing"
)

var pageListServiceConfigFile = "../secrets/config.blogger-tools-lib-test.json"

func TestPageListService_Get(t *testing.T) {
	config, err := blogger_tools_lib.NewConfig(pageListServiceConfigFile)
	if err != nil {
		t.Error(err)
		return
	}

	blogger := blogger_tools_blogger.NewBlogger(nil, *config.Blogger.Blogs[0].AccessToken, config.Blogger.Blogs[0].ID)
	pageList, err := blogger.PageList.Get()
	if err != nil {
		t.Error(err)
		return
	}

	if pageList.Kind != "blogger#pageList" {
		t.Errorf("expected blogger#pageList, got %v", pageList.Kind)
	}

	if pageList.TotalItems != 1 {
		t.Errorf("expected 1, got %v", pageList.TotalItems)
		return
	}
}
