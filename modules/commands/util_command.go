package utilcommand

import (
	"github.com/urfave/cli/v2"
	services "github.com/yukkyun/eccu/modules/services"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "out",
				Aliases: []string{"o"},
				Usage:   "choose output type (default: tsv)",
			},
			&cli.StringFlag{
				Name:        "status",
				Aliases:     []string{"s"},
				Usage:       "choose ec2 status (default: return all instances)",
				DefaultText: "all",
			},
		},
		Action: services.GetEc2List,
	}
}

func FuzzySearchCommand() *cli.Command {
	return &cli.Command{
		Name:    "fs",
		Aliases: []string{"f"},
		Usage:   "Do Fuzzy-Search against AWS resources. Target resource is given by subcommand.",
		Subcommands: []*cli.Command{
			{
				Name:  "ec2",
				Usage: "Do Fuzzy-Search against EC2",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "status",
						Aliases: []string{"s"},
						Usage:   "choose ec2 status (default: return all instances)",
					},
				},
				Action: services.EC2FuzzySearch,
			},
		},
	}
}
