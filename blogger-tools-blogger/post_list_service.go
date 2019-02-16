package blogger_tools_blogger

type PostListService struct {
	*Client
}

func (service *PostListService) Get() (postList *PostList, err error) {
	body, err := service.SendRequest("/posts")

	postList, err = NewPostListFromJSON(body)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return postList, nil
}
