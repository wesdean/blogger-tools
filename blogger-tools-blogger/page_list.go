package blogger_tools_blogger

import (
	"encoding/json"
	"errors"
)

type PageList struct {
	TotalItems int    `json:"totalItems,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Kind       string `json:"kind,omitempty"`
	ETag       string `json:"etag,omitempty"`
	Items      []Page
}

func NewPageListFromJSON(data []byte) (pageList *PageList, err error) {
	pageList = &PageList{}
	return pageList, pageList.parseJSON(data)
}

func (pageList *PageList) parseJSON(data []byte) error {
	if pageList == nil {
		return errors.New("Cannot unmarshal into nil")
	}

	err := json.Unmarshal(data, pageList)
	if err != nil {
		return err
	}

	if pageList.Items != nil {
		pageList.TotalItems = len(pageList.Items)
	}

	return nil
}
