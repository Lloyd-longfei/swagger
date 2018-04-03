package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
)

func main(){
	app :=cli.NewApp()

	app.Name = "hello"

	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name : "language",
			Aliases: []string{"lang"},
			Usage: "change language",

			Action: func(c *cli.Context) error{
				language := c.Args().First()
				if language == "chinese"{
					fmt.Println("language is 中文")
				}else {
					fmt.Println("language is english")
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}