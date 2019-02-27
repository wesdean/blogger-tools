package blogger_tools_lib_oauth_tool

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/google/oauth2l/sgauth"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"io/ioutil"
	"os"
	"strings"
)

type OAuthTool struct {
	*blogger_tools_lib.BloggerTool
}

type OAuthToolArgs struct {
	ResetLog        bool
	AuthFlowHandler string
	BlogIds         []string
}

type OAuthToolResults struct {
	Success bool
	Blogs   map[string]string
}

func (tool *OAuthTool) Run(args *OAuthToolArgs) (results *OAuthToolResults, err error) {
	results = &OAuthToolResults{Success: false, Blogs: map[string]string{}}

	log, err := tool.Config.CreateLogger(tool.Config.BuildLogFilePath(tool.Config.Logs.OAuthTool), args.ResetLog)
	if err != nil {
		return results, err
	}
	defer log.Close()

	log.Info("Running OAuthTool")

	errorFlag := false
	for index, blogConfig := range tool.Config.Blogger.Blogs {
		if len(args.BlogIds) > 0 {
			if !tool.hasBlogId(blogConfig.ID, args.BlogIds) {
				log.Infof("%s skipped: not selected", blogConfig.ID)
				results.Blogs[blogConfig.ID] = "skipped: not selected"
				continue
			}
		}

		if *blogConfig.OAuthKeyFile == "" {
			log.Infof("%s skipped: no key file", blogConfig.ID)
			results.Blogs[blogConfig.ID] = "skipped: no key file"
			continue
		}

		log.Infof("%s running", blogConfig.ID)
		message, err := tool.RunForBlog(&tool.Config.Blogger.Blogs[index], args.AuthFlowHandler)
		if err != nil {
			log.Error(strings.Replace(message, "\n", "", -1))
			results.Blogs[blogConfig.ID] = err.Error()
			errorFlag = true
			continue
		}

	}

	if errorFlag {
		log.Error("authentication errors occurred")
		return results, errors.New("authentication errors occurred")
	} else {
		log.Info("OAuthTool successful")
		results.Success = true
	}

	return results, err
}

func (tool *OAuthTool) GetFlowHandler(name string) (func(url string) (string, error), error) {
	switch name {
	case "cli":
		return func(url string) (string, error) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Visit this URL get an access token:\n%s\n\n", url)
			fmt.Println("Enter access token")
			fmt.Print("-> ")
			code, _ := reader.ReadString('\n')
			code = strings.Replace(code, "\r\n", "", -1)
			code = strings.Replace(code, "\n", "", -1)

			return code, nil
		}, nil
	default:
		return nil, errors.New("flow handler does not exist")
	}
}

func (tool *OAuthTool) RunForBlog(blogConfig *blogger_tools_lib.BlogConfig, flowHandlerName string) (string, error) {
	flowHandler, err := tool.GetFlowHandler(flowHandlerName)
	if err != nil {
		return err.Error(), err
	}

	authKeyFile, err := ioutil.ReadFile(*blogConfig.OAuthKeyFile)
	if err != nil {
		return fmt.Sprintf("failed: %s", err.Error()), err
	}

	settings := &sgauth.Settings{
		CredentialsJSON:  string(authKeyFile),
		Scope:            "https://www.googleapis.com/auth/blogger",
		OAuthFlowHandler: flowHandler,
		State:            "state",
	}

	token, err := sgauth.FetchToken(context.Background(), settings)
	if err != nil {
		return fmt.Sprintf("failed: %s", strings.Replace(err.Error(), "\n", "", -1)), err
	}

	blogConfig.AccessToken = &token.AccessToken

	if blogConfig.AccessTokenFile != nil {
		err = ioutil.WriteFile(tool.Config.BuildSecretFilePath(*blogConfig.AccessTokenFile), []byte(token.AccessToken), 0666)
		if err != nil {
			return fmt.Sprintf("success: %s", err), err
		}
	}

	return "success", nil
}

func (tool *OAuthTool) hasBlogId(blogId string, blogIds []string) bool {
	for _, id := range blogIds {
		if blogId == id {
			return true
		}
	}
	return false
}
