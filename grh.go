package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Github Remote URL Generator"
	app.Usage = "Making working with multiple github accounts hassle free"
	app.Description = "Generate git remote origin command when using multiple github accounts on one workstation."
	app.Author = "Tremaine Buchanan"
	app.Version = "1.0.0"
}

func flags() {
	var orgname string
	var username string
	var reponame string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "orgname, o",
			Usage:       "name of organization",
			Destination: &orgname,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "username, u",
			Usage:       "github user name tied to the organization",
			Destination: &username,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "reponame, r",
			Usage:       "name of the repository",
			Destination: &reponame,
			Required:    true,
		},
	}

	app.Action = func(c *cli.Context) error {
		cmd := "git remote add origin git@github-"
		result := cmd + orgname + ":" + username + "/" + reponame + ".git"
		fmt.Println(result)
		return nil
	}
}

func main() {
	info()
	flags()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
