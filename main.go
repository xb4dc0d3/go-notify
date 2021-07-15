package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func main() {
	app := &cli.App{
		Name:        "Go-Notify",
		HelpName:    "",
		Usage:       "A fsnotify wrapper using tmux as intermediary",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "0.1.0",
		Description: "A fsnotify wrapper using tmux as intermediary",
		Commands: []*cli.Command{
			{
				Name:      "attach",
				Usage:     "Attach the folder or file you want to watch",
				UsageText: "go-notify attach <file_path>|<folder_path>",
				Action: func(c *cli.Context) error {

					if c.NArg() < 0 || c.Args().Get(0) == "" {
						return cli.ShowCommandHelp(c, "attach")
					}

					notify_path := c.Args().First()
					GoNotify(notify_path)

					return nil
				},
			},
		},
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete: func(*cli.Context) {
		},
		CommandNotFound: func(*cli.Context, string) {
		},
		Compiled:  time.Time{},
		Authors:   []*cli.Author{{Name: "xb4dc0d3", Email: ""}},
		Copyright: "",
		Reader:    nil,
		Writer:    nil,
		ErrWriter: nil,
		ExitErrHandler: func(context *cli.Context, err error) {
		},
		Metadata:               map[string]interface{}{},
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
