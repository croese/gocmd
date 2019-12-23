package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "v1.0.0"

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version", Aliases: []string{"V"},
		Usage: "print the version",
	}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "b",
				Usage: "Number the non-blank output lines, starting at 1",
			},
			&cli.BoolFlag{
				Name:  "e",
				Usage: "Display non-printing characters (see the -v option), and display a dollar sign ('$') at the end of each line",
			},
			&cli.BoolFlag{
				Name:  "n",
				Usage: "Number the output lines, starting at 1",
			},
			&cli.BoolFlag{
				Name:  "s",
				Usage: "Squeeze multiple adjacent empty lines, causing the output to be single spaced",
			},
			&cli.BoolFlag{
				Name:  "t",
				Usage: "Display non-printing characters (see the -v option), and display tab characters as '^I'",
			},
			&cli.BoolFlag{
				Name:  "u",
				Usage: "Disable output buffering",
			},
			&cli.BoolFlag{
				Name:  "v",
				Usage: "Display non-printing characters so they are visible.  Control characters print as '^X' for control-X; the delete character (octal 0177) prints as '^?'.  Non-ASCII characters (with the high bit set) are printed as 'M-' (for meta) followed by the character for the low 7 bits",
			},
		},
		Name:    "cat",
		Version: version,
		Usage:   "concatenate and print files",
		Action: func(c *cli.Context) error {
			fmt.Println("hello world")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
	}
}
