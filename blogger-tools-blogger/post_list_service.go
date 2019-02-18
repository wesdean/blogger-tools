package blogger_tools_blogger

import "github.com/kr/pretty"

type PostListService struct {
	*Client
}

func (service *PostListService) Get() (postList *PostList, err error) {
	body, err := service.SendRequest("/posts")
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	pretty.Println(string(body))
	postList, err = NewPostListFromJSON(body)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return postList, nil
}
