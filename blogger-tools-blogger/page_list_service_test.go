package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
)

func TestPageListService_Get(t *testing.T) {
	config := getConfig()

	blogger := blogger_tools_blogger.NewBlogger(nil, config.APIKey, config.BlogIDs[0])
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
