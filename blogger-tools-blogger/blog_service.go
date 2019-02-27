package blogger_tools_blogger

import (
	"encoding/json"
)

type BlogService struct {
	*Client
}

/*
https://developers.google.com/blogger/docs/3.0/reference/blogs/get
*/
func (service *BlogService) Get() (blog *Blog, err error, response *ErrorResponse) {
	body, err, response := service.SendRequest("/", nil)
	if err != nil {
		service.logger.Error(err)
		return nil, err, response
	}

	blog = &Blog{}
	err = json.Unmarshal(body, blog)
	if err != nil {
		service.logger.Error(err)
		return nil, err, response
	}

	return blog, nil, response
}
