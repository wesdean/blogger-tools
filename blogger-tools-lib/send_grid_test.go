package blogger_tools_lib_test

import (
	"encoding/json"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"github.com/wesdean/blogger-tools/blogger-tools-lib/blogger-tools-lib-notify-tool"
	"io/ioutil"
	"testing"
)

func TestSendGrid_SendMail(t *testing.T) {
	config, err := blogger_tools_lib.NewConfig(bloggerToolConfigFile)
	if err != nil {
		t.Error(err)
		return
	}

	recipientFile := config.BuildSecretFilePath(config.NotifyTool.BlogUpdatedRecipientsFile)
	recipientJSON, err := ioutil.ReadFile(recipientFile)
	if err != nil {
		t.Error(err)
		return
	}

	var allRecipients map[string][]blogger_tools_lib_notify_tool.BlogUpdatedRecipient
	err = json.Unmarshal(recipientJSON, &allRecipients)
	if err != nil {
		t.Error(err)
		return
	}

	var recipients []*mail.Email
	for _, value := range allRecipients["3051261493420306591"] {
		recipients = append(recipients, mail.NewEmail(value.Name, value.Email))
	}

	sendGrid := blogger_tools_lib.NewSendGrid(config.SendGrid.APIKey)
	_, err = sendGrid.SendMail(
		mail.NewEmail(config.SendGrid.DefaultFromName, config.SendGrid.DefaultFromEmail),
		recipients,
		"TestSendGrid_SendMail Test",
		mail.NewContent("text/html", "This is a test for %function%"),
		&map[string]string{"%function%": "TestSendGrid_SendMail"},
	)
}
