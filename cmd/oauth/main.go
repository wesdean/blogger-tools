package main

import (
	"encoding/json"
	"errors"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"github.com/wesdean/blogger-tools/blogger-tools-lib/blogger-tools-lib-oauth-tool"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "OAuth Tool"
	app.Version = "0.1.0"
	app.Usage = "perform oauth workflow to gain access to the Google Blogger API"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "",
			Usage: "filename of the configuration file to use.",
		},
		cli.StringFlag{
			Name:  "blogs",
			Value: "",
			Usage: "ids of the blogs to authenticate.",
		},
	}

	var config *blogger_tools_lib.Config
	var tool *blogger_tools_lib_oauth_tool.OAuthTool
	var blogIds = []string{}

	app.Before = func(context *cli.Context) error {
		var err error

		configFilename := context.GlobalString("config")
		if configFilename == "" {
			configFilename = "./conf.json"
		}
		config, err = blogger_tools_lib.NewConfig(configFilename)
		if err != nil {
			return err
		}

		bloggerTool := &blogger_tools_lib.BloggerTool{Config: config}
		tool = &blogger_tools_lib_oauth_tool.OAuthTool{bloggerTool}

		for _, id := range strings.Split(context.String("blogs"), ",") {
			id = strings.TrimSpace(id)
			if id != "" {
				blogIds = append(blogIds, id)
			}
		}

		return err
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run the oauth workflow",
			Action: func(context *cli.Context) error {
				results, err := tool.Run(&blogger_tools_lib_oauth_tool.OAuthToolArgs{
					AuthFlowHandler: "cli",
					BlogIds:         blogIds,
				})
				if err != nil {
					return err
				}

				if !results.Success {
					jsonBytes, err := json.Marshal(results)
					if err != nil {
						return err
					}
					return errors.New(string(jsonBytes))
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
