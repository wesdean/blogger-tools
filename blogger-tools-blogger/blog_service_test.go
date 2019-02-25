package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"testing"
)

var blogServiceConfigFile = "../secrets/config.blogger-tools-lib-test.json"

func TestBlogService_Get(t *testing.T) {
	config, err := blogger_tools_lib.NewConfig(blogServiceConfigFile)
	if err != nil {
		t.Error(err)
		return
	}

	blogger := blogger_tools_blogger.NewBlogger(nil, *config.Blogger.Blogs[0].AccessToken, config.Blogger.Blogs[0].ID)
	blog, err, response := blogger.Blog.Get()
	if err != nil {
		t.Error(err)
		return
	}
	if response != nil {
		t.Error(response)
		return
	}
	if blog.Id != "3051261493420306591" {
		t.Errorf("expected 3051261493420306591, got %v", blog.Id)
		return
	}
}
