package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "v1.0.0"

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "n",
				Usage: "do not print the trailing newline character",
			},
		},
		Name:    "echo",
		Version: version,
		Usage:   "write arguments to the standard output",
		Action: func(c *cli.Context) error {
			skipNewline := c.Bool("n")
			for i, arg := range c.Args().Slice() {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(arg)
			}

			if !skipNewline {
				fmt.Println()
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
	}
}
