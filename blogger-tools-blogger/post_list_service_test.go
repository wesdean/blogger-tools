package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
)

func TestPostListService_Get(t *testing.T) {
	config := getConfig()

	blogger := blogger_tools_blogger.NewBlogger(nil, config.APIKey, config.BlogIDs[0])
	postList, err := blogger.PostList.Get()
	if err != nil {
		t.Error(err)
		return
	}

	if postList.Kind != "blogger#postList" {
		t.Errorf("expected blogger#postList, got %v", postList.Kind)
	}

	if postList.TotalItems != 1 {
		t.Errorf("expected 1, got %v", postList.TotalItems)
		return
	}
}
