package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
)

func TestBlogService_Get(t *testing.T) {
	config := getConfig()

	blogger := blogger_tools_blogger.NewBlogger(nil, config.APIKey, config.BlogIDs[0])
	blog, err := blogger.Blog.Get()
	if err != nil {
		t.Error(err)
		return
	}
	if blog.Id != "3960547499512363533" {
		t.Errorf("expected 3960547499512363533, got %v", blog.Id)
		return
	}
}
