package blogger_tools_blogger

import (
	"encoding/json"
	"errors"
	"time"
)

type Post struct {
	Kind      string       `json:"kind,omitempty"`
	Id        string       `json:"id,omitempty"`
	Blog      *Blog        `json:"blog,omitempty"`
	Published time.Time    `json:"published,omitempty"`
	Updated   time.Time    `json:"updated,omitempty"`
	ETag      string       `json:"etag,omitempty"`
	URL       string       `json:"url,omitempty"`
	SelfLink  string       `json:"selfLink,omitempty"`
	Title     string       `json:"title,omitempty"`
	Content   string       `json:"content,omitempty"`
	Author    *Author      `json:"author,omitempty"`
	Replies   *RepliesList `json:"replies,omitempty"`
}

type RepliesList struct {
	TotalItems int    `json:"totalItems,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
}

func NewPostFromJSON(data []byte) (post *Post, err error) {
	post = &Post{}
	return post, post.parseJSON(data)
}

func (post *Post) parseJSON(data []byte) error {
	if post == nil {
		return errors.New("Cannot unmarshal into nil")
	}

	return json.Unmarshal(data, post)
}
