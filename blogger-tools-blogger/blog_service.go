package blogger_tools_blogger

import (
	"encoding/json"
)

type BlogService struct {
	*Client
}

func (service *BlogService) Get() (blog *Blog, err error) {
	body, err, _ := service.SendRequest("/", nil)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	blog = &Blog{}
	err = json.Unmarshal(body, blog)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return blog, nil
}
