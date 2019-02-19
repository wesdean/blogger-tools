package blogger_tools_blogger

import (
	"strconv"
	"strings"
	"time"
)

type PostListService struct {
	*Client
}

/*
https://developers.google.com/blogger/docs/3.0/reference/posts/list#parameters
*/
type PostListServiceOptions struct {
	useEndDate bool
	endDate    time.Time

	useStartDate bool
	startDate    time.Time

	useFetchBodies bool
	fetchBodies    bool

	useFetchImages bool
	fetchImages    bool

	useLabels bool
	labels    []string

	useMaxResults bool
	maxResults    int

	useOrderBy bool
	orderBy    string

	usePageToken bool
	pageToken    string

	useStatus bool
	status    string

	useView bool
	view    string
}

func NewPostListServiceOptions() *PostListServiceOptions {
	return &PostListServiceOptions{}
}

func (options *PostListServiceOptions) EndDate(endDate time.Time) *PostListServiceOptions {
	options.useEndDate = true
	options.endDate = endDate
	return options
}

func (options *PostListServiceOptions) StartDate(startDate time.Time) *PostListServiceOptions {
	options.useStartDate = true
	options.startDate = startDate
	return options
}

func (options *PostListServiceOptions) FetchBodies(fetchBodies bool) *PostListServiceOptions {
	options.useFetchBodies = true
	options.fetchBodies = fetchBodies
	return options
}

func (options *PostListServiceOptions) FetchImages(fetchImages bool) *PostListServiceOptions {
	options.useFetchImages = true
	options.fetchImages = fetchImages
	return options
}

func (options *PostListServiceOptions) MaxResults(max int) *PostListServiceOptions {
	options.useMaxResults = true
	options.maxResults = max
	return options
}

func (options *PostListServiceOptions) OrderBy(orderBy string) *PostListServiceOptions {
	options.useOrderBy = true
	options.orderBy = orderBy
	return options
}

func (options *PostListServiceOptions) PageToken(token string) *PostListServiceOptions {
	options.usePageToken = true
	options.pageToken = token
	return options
}

func (options *PostListServiceOptions) Status(status string) *PostListServiceOptions {
	options.useStatus = true
	options.status = status
	return options
}

func (options *PostListServiceOptions) View(view string) *PostListServiceOptions {
	options.useView = true
	options.view = view
	return options
}

func (options *PostListServiceOptions) Labels(labels ...string) *PostListServiceOptions {
	options.useLabels = true
	options.labels = labels
	return options
}

/*
https://developers.google.com/blogger/docs/3.0/reference/posts/list#http-request
*/
func (service *PostListService) Get(options *PostListServiceOptions) (postList *PostList, err error) {
	var params map[string]string

	if options != nil {
		params = map[string]string{}

		if options.useEndDate {
			params["endDate"] = options.endDate.Format(time.RFC3339)
		}
		if options.useStartDate {
			params["startDate"] = options.startDate.Format(time.RFC3339)
		}
		if options.useFetchBodies {
			params["fetchBodies"] = strconv.FormatBool(options.fetchBodies)
		}
		if options.useFetchImages {
			params["fetchImages"] = strconv.FormatBool(options.fetchImages)
		}
		if options.useMaxResults {
			params["maxResults"] = strconv.Itoa(options.maxResults)
		}
		if options.useOrderBy {
			params["orderBy"] = options.orderBy
		}
		if options.usePageToken {
			params["pageToken"] = options.pageToken
		}
		if options.useStatus {
			params["status"] = options.status
		}
		if options.useView {
			params["view"] = options.view
		}
		if options.useLabels {
			params["labels"] = strings.Join(options.labels[:], ",")
		}
	}

	body, err := service.SendRequest("/posts", params)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	postList, err = NewPostListFromJSON(body)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return postList, nil
}
