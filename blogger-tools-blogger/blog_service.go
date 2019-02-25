package blogger_tools_blogger

import (
	"encoding/json"
)

type BlogService struct {
	*Client
}

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
