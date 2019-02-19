package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
	"time"
)

func TestPostListService_Get(t *testing.T) {
	//if skipExternalServices {
	//	t.Skip("Uses external services")
	//	return
	//}

	config, err := getConfig()
	if err != nil {
		t.Error(err)
		return
	}

	blogger := blogger_tools_blogger.NewBlogger(getLogger(), config.Blogs[0].AccessToken, config.Blogs[0].ID)

	t.Run("Fetch all posts", func(t *testing.T) {
		postList, err := blogger.PostList.Get(nil)
		if err != nil {
			t.Error(err)
			return
		}

		if postList.Kind != "blogger#postList" {
			t.Errorf("expected blogger#postList, got %v", postList.Kind)
		}

		if postList.TotalItems != 2 {
			t.Errorf("expected 2, got %v", postList.TotalItems)
			return
		}
	})

	t.Run("Fetch latest post", func(t *testing.T) {
		postListOptions := blogger_tools_blogger.NewPostListServiceOptions().
			OrderBy("published").
			MaxResults(1)

		postList, err := blogger.PostList.Get(postListOptions)
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

		publishedString := "2019-02-18T14:47:00-08:00"
		if postList.Items[0].Published.Format(time.RFC3339) != publishedString {
			t.Errorf("expected %v, got %v", publishedString, postList.Items[0].Published.Format(time.RFC3339))
		}
	})
}
