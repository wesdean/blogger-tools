package blogger_tools_blogger

import "github.com/google/logger"

type Blogger struct {
	Blog     *BlogService
	PageList *PageListService
}

func NewBlogger(logger *logger.Logger, apiKey string, blogId string) *Blogger {
	client := NewClient(logger, apiKey, blogId)
	return &Blogger{
		Blog:     &BlogService{Client: client},
		PageList: &PageListService{Client: client},
	}
}
