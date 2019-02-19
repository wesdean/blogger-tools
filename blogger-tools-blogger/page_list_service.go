package blogger_tools_blogger

type PageListService struct {
	*Client
}

func (service *PageListService) Get() (pageList *PageList, err error) {
	body, err := service.SendRequest("/pages", nil)

	pageList, err = NewPageListFromJSON(body)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return pageList, nil
}
