package blogger_tools_blogger

import (
	"encoding/json"
	"errors"
)

type PostList struct {
	TotalItems int    `json:"totalItems,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Kind       string `json:"kind,omitempty"`
	ETag       string `json:"etag,omitempty"`
	Items      []Page
}

func NewPostListFromJSON(data []byte) (postList *PostList, err error) {
	postList = &PostList{}
	return postList, postList.parseJSON(data)
}

func (postList *PostList) parseJSON(data []byte) error {
	if postList == nil {
		return errors.New("Cannot unmarshal into nil")
	}

	err := json.Unmarshal(data, postList)
	if err != nil {
		return err
	}

	if postList.Items != nil {
		postList.TotalItems = len(postList.Items)
	}

	return nil
}
