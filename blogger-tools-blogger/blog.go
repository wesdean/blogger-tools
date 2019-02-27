package blogger_tools_blogger

import (
	"encoding/json"
	"errors"
	"time"
)

/*
https://developers.google.com/blogger/docs/3.0/reference/blogs
*/
type Blog struct {
	Kind        string    `json:"kind,omitempty"`
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Published   time.Time `json:"published,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
	URL         string    `json:"url,omitempty"`
	SelfLink    string    `json:"selfLink,omitempty"`
	Posts       PostList  `json:"posts,omitempty"`
	Pages       PageList  `json:"pages,omitempty"`
	Locale      Locale    `json:"locale,omitempty"`
}

type Locale struct {
	Language string `json:"language,omitempty"`
	Country  string `json:"country,omitempty"`
	Variant  string `json:"variant,omitempty"`
}

type Author struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	URL         string `json:"url,omitempty"`
	Image       *Image `json:"image,omitempty"`
}

type Image struct {
	URL string `json:"url,omitempty"`
}

func NewBlogFromJSON(data []byte) (blog *Blog, err error) {
	blog = &Blog{}
	return blog, blog.parseJSON(data)
}

func (blog *Blog) parseJSON(data []byte) error {
	if blog == nil {
		return errors.New("Cannot unmarshal into nil")
	}

	return json.Unmarshal(data, blog)
}
