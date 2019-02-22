package blogger_tools_lib

import (
	"errors"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	APIKey   string
	Endpoint string
	Host     string
}

func NewSendGrid(apiKey string) *SendGrid {
	return &SendGrid{
		APIKey:   apiKey,
		Endpoint: "/v3",
		Host:     "https://api.sendgrid.com",
	}
}

func (sendGrid *SendGrid) SendMail(
	from *mail.Email,
	tos []*mail.Email,
	subject string,
	content *mail.Content,
	substitutions *map[string]string,
) (*rest.Response, error) {
	newMail := mail.NewV3Mail()
	newMail.SetFrom(from)
	newMail.AddContent(content)

	personalization := mail.NewPersonalization()
	personalization.AddTos(tos...)
	if substitutions != nil {
		for key, value := range *substitutions {
			personalization.SetSubstitution(key, value)
		}
	}
	personalization.Subject = subject
	newMail.AddPersonalizations(personalization)

	request := sendgrid.GetRequest(
		sendGrid.APIKey,
		sendGrid.Endpoint+"/mail/send",
		sendGrid.Host)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(newMail)
	response, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 202 {
		return nil, errors.New(response.Body)
	}

	return response, nil
}
