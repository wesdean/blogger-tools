package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
)

func TestBlogService_Get(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error(err)
		return
	}

	blogger := blogger_tools_blogger.NewBlogger(nil, config.Blogs[0].AccessToken, config.Blogs[0].ID)
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
