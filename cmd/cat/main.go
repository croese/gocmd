package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "v1.0.0"

type args struct {
	numberNonBlank     bool
	displayLineEnds    bool
	numberLines        bool
	squeezeLines       bool
	displayTabs        bool
	disableBuffering   bool
	displayNonPrinting bool
}

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
		Name:        "cat",
		Version:     version,
		Usage:       "concatenate and print files",
		UsageText:   "cat [-benstuv] [file ...]",
		Description: "The cat utility reads files sequentially, writing them to the standard output.  The file operands are processed in command-line order.  If file is a single dash ('-') or absent, cat reads from the standard input.",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				handleStdin()
			} else {
				args := args{
					numberNonBlank:     c.Bool("b"),
					displayLineEnds:    c.Bool("e"),
					numberLines:        c.Bool("n"),
					squeezeLines:       c.Bool("s"),
					displayTabs:        c.Bool("t"),
					disableBuffering:   c.Bool("u"),
					displayNonPrinting: c.Bool("v"),
				}
				for _, file := range c.Args().Slice() {
					if file == "-" {
						handleStdin(args)
					} else {
						handleFile(file, args)
					}
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
	}
}

func handleStdin(a args) {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		fmt.Println(stdin.Text())
	}
}

func handleFile(filename string, a args) {
	fmt.Println("FILE:", filename)
}
