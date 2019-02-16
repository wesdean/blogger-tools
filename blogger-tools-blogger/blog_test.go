package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
)

func TestNewBlogFromJSON(t *testing.T) {
	blog, err := blogger_tools_blogger.NewBlogFromJSON([]byte(blogExampleJSON))
	if err != nil {
		t.Error(err)
		return
	}

	if blog.Kind != "blogger#blog" {
		t.Errorf("expected blogger#blog, got %v", blog.Kind)
		return
	}

	if blog.Id != "2399953" {
		t.Errorf("expected 2399953, got %v", blog.Id)
		return
	}
}

var blogExampleJSON = `{
  "kind": "blogger#blog",
  "id": "2399953",
  "name": "Blogger Buzz",
  "description": "The Official Buzz from Blogger at Google",
  "published": "2007-04-23T22:17:29.261Z",
  "updated": "2011-08-02T06:01:15.941Z",
  "url": "http://buzz.blogger.com/",
  "selfLink": "https://www.googleapis.com/blogger/v3/blogs/2399953",
  "posts": {
    "totalItems": 494,
    "selfLink": "https://www.googleapis.com/blogger/v3/blogs/2399953/posts"
  },
  "pages": {
    "totalItems": 2,
    "selfLink": "https://www.googleapis.com/blogger/v3/blogs/2399953/pages"
  },
  "locale": {
    "language": "en",
    "country": "",
    "variant": ""
  }
}`
