package main

import (
	"fmt"
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"github.com/wesdean/blogger-tools/blogger-tools-lib/blogger-tools-lib-notify-tool"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Blogger Notify Tool"
	app.Version = "0.1.0"
	app.Usage = "send notification of blog events"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "",
			Usage: "filename of the configuration file to use.",
		},
	}

	var config *blogger_tools_lib.Config
	var tool *blogger_tools_lib_notify_tool.NotifyTool

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
		tool = &blogger_tools_lib_notify_tool.NotifyTool{bloggerTool}

		return err
	}

	app.Commands = []cli.Command{
		{
			Name:  "diagnostic",
			Usage: "run notify tool diagnostics",
			Action: func(context *cli.Context) error {
				results, err := tool.Run(&blogger_tools_lib_notify_tool.NotifyToolArgs{
					ResetLog:         true,
					OAuthFlowHandler: "cli",
				})
				if err != nil {
					return err
				}

				if results.Success != true {
					return err
				}

				fmt.Println("Result: Pass")
				return nil
			},
		},
		{
			Name:  "blog-update",
			Usage: "send notification for a blog update",
			Action: func(context *cli.Context) error {
				results, err := tool.Run(&blogger_tools_lib_notify_tool.NotifyToolArgs{
					ResetLog:         true,
					OAuthFlowHandler: "cli",
					Actions: &blogger_tools_lib_notify_tool.NotifyToolActions{
						&blogger_tools_lib_notify_tool.ActionBlogUpdatedOptions{},
					},
				})
				if err != nil {
					return err
				}

				if results.Success != true {
					return err
				}

				fmt.Println("done")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
