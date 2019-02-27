package blogger_tools_blogger

type PageListService struct {
	*Client
}

/*
https://developers.google.com/blogger/docs/3.0/reference/pages/get
*/
func (service *PageListService) Get() (pageList *PageList, err error) {
	body, err, _ := service.SendRequest("/pages", nil)

	pageList, err = NewPageListFromJSON(body)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return pageList, nil
}
