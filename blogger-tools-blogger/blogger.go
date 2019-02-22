package blogger_tools_blogger

import "github.com/google/logger"

type Blogger struct {
	Blog     *BlogService
	PageList *PageListService
	PostList *PostListService
}

func NewBlogger(logger *logger.Logger, accessToken string, blogId string) *Blogger {
	client := NewClient(logger, accessToken, blogId)
	return &Blogger{
		Blog:     &BlogService{Client: client},
		PageList: &PageListService{Client: client},
		PostList: &PostListService{Client: client},
	}
}
