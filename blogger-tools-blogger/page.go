package blogger_tools_blogger

import (
	"encoding/json"
	"errors"
	"time"
)

/*
https://developers.google.com/blogger/docs/3.0/reference/pages
*/
type Page struct {
	Kind      string    `json:"kind,omitempty"`
	Id        string    `json:"id,omitempty"`
	Blog      *Blog     `json:"blog,omitempty"`
	Published time.Time `json:"published,omitempty"`
	Updated   time.Time `json:"updated,omitempty"`
	ETag      string    `json:"etag,omitempty"`
	URL       string    `json:"url,omitempty"`
	SelfLink  string    `json:"selfLink,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Author    *Author   `json:"author,omitempty"`
}

func NewPageFromJSON(data []byte) (page *Page, err error) {
	page = &Page{}
	return page, page.parseJSON(data)
}

func (page *Page) parseJSON(data []byte) error {
	if page == nil {
		return errors.New("Cannot unmarshal into nil")
	}

	return json.Unmarshal(data, page)
}
