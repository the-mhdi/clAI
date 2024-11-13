package main

import (
	"fmt"
	"log"
	"os"

	"github.com/the-mhdi/maShit/util"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "config",
				Usage: "load the config",
				Action: func(cCtx *cli.Context) error {
					config, err := util.LoadConfig(cCtx.Args().Get(0))
					fmt.Printf("config loaded %s\n", config) //test
					loadModel(config)

					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
