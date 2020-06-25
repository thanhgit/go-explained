package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)
func main() {
	app := cli.NewApp()
	app.Name = "hello-cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
		cli.StringFlag{
			Name:        "user, u",
			Usage:       "Who am i",
			Value:       "root",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)
		return nil
	}
	app.Action = func(c *cli.Context) error {
		user := c.GlobalString("user")
		fmt.Println("My name is " + user)
		return nil
	}
	app.Run(os.Args)
}
