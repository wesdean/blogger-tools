package blogger_tools_lib_notify_tool

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/wesdean/blogger-tools/blogger-tools-blogger"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
)

type BlogUpdatedAction struct {
	*Action
	Recipients []BlogUpdatedRecipient
}

type BlogUpdatedRecipient struct {
	Name  string
	Email string
}

func (action *BlogUpdatedAction) Do() error {
	post, err := action.getMostRecentPost()
	if err != nil {
		return err
	}

	if post != nil {
		recipients := []*mail.Email{}
		for _, recipient := range action.Recipients {
			recipients = append(recipients, mail.NewEmail(recipient.Name, recipient.Email))
		}
		if len(recipients) > 0 {
			content := mail.NewContent(
				"text/html",
				"%blogName% has an update! Check it out!\n\n%blogLink%",
			)
			substitutions := &map[string]string{
				"%blogName%": action.Blog.Name,
				"%blogLink%": action.Blog.URL,
			}
			subject := fmt.Sprintf("%s has an update! Check it out!", action.Blog.Name)

			sendGrid := blogger_tools_lib.NewSendGrid(action.Config.SendGrid.APIKey)
			_, err := sendGrid.SendMail(
				mail.NewEmail(action.Config.SendGrid.DefaultFromName, action.Config.SendGrid.DefaultFromEmail),
				recipients,
				subject,
				content,
				substitutions,
			)
			if err != nil {
				action.Logger.Error(err)
				return err
			}
		}
	}

	return nil
}

func (action *BlogUpdatedAction) getMostRecentPost() (*blogger_tools_blogger.Post, error) {
	blogger := blogger_tools_blogger.NewBlogger(action.Logger, action.BlogAccessToken, action.Blog.Id)
	postListOptions := blogger_tools_blogger.NewPostListServiceOptions().MaxResults(1).OrderBy("published")
	postList, err := blogger.PostList.Get(postListOptions)
	if err != nil {
		action.Logger.Error(err)
		return nil, err
	}

	if postList.TotalItems > 0 {
		return &postList.Items[0], nil
	} else {
		action.Logger.Warningf("No posts found in blog: %s", action.Blog.Id)
		return nil, nil
	}
}
