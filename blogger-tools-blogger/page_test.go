package blogger_tools_blogger_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"testing"
)

func TestNewPageFromJSON(t *testing.T) {
	page, err := blogger_tools_blogger.NewPageFromJSON([]byte(pageExampleJSON))
	if err != nil {
		t.Error(err)
		return
	}

	if page.Kind != "blogger#page" {
		t.Errorf("expected blogger#page, got %v", page.Kind)
	}
}

var pageExampleJSON = `{
	"kind": "blogger#page",
	"id": "7706273476706534553",
	"blog": {
		"id": "2399953"
	},
	"published": "2011-08-01T19:58:00.000Z",
	"updated": "2011-08-01T19:58:51.947Z",
	"url": "http://buzz.blogger.com/2011/08/latest-updates-august-1st.html",
	"selfLink": "https://www.googleapis.com/blogger/v3/blogs/2399953/pages/7706273476706534553",
	"title": "Latest updates, August 1st",
	"content": "elided for readability",
	"author": {
		"id": "401465483996",
		"displayName": "Brett Wiltshire",
		"url": "http://www.blogger.com/profile/01430672582309320414",
		"image": {
			"url": "http://4.bp.blogspot.com/_YA50adQ-7vQ/S1gfR_6ufpI/AAAAAAAAAAk/1ErJGgRWZDg/S45/brett.png"
		}
	}
}`
