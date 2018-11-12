package main

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	commandGen,
	//commandDiff,
}

var commandGen = cli.Command{
	Name:  "gen",
	Usage: "Generate test code",
	Description: `
    Generete test code according to yaml and template.
`,
	Action: doGen,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "yaml",
			Value: "tests.yml",
			Usage: "yaml file defines requests and responses for testing",
		},
		cli.StringFlag{
			Name:  "template",
			Value: "template_test.go",
			Usage: "template file defines test code",
		},
	},
}

func doGen(c *cli.Context) error {
	return nil
}
