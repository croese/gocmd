package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "v1.0.0"

const description = `The wc utility displays the number of lines, words, and bytes contained in each input file, 
or standard input (if no file is specified) to the standard output.  A line is defined as a string of characters 
delimited by a <newline> character.  Characters beyond the final <newline> character will not be included in the 
line count.
		   
A word is defined as a string of characters delimited by white space characters.  White space characters are the 
set of characters for which the 'unicode.IsSpace' function returns true.  If more than one input file is specified, a 
line of cumulative counts for all the files is displayed on a separate line after the output for the last file.`

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "c",
				Usage: "The number of bytes in each input file is written to the standard output.  This will cancel out any prior usage of the -m option.",
			},
			&cli.BoolFlag{
				Name:  "l",
				Usage: "The number of lines in each input file is written to the standard output.",
			},
			&cli.BoolFlag{
				Name:  "m",
				Usage: "The number of characters in each input file is written to the standard output.  If the current locale does not support multibyte characters, this is equivalent to the -c option.  This will cancel out any prior usage of the -c option.",
			},
			&cli.BoolFlag{
				Name:  "w",
				Usage: "The number of words in each input file is written to the standard output.",
			},
		},
		Name:        "wc",
		Version:     version,
		Usage:       "word, line, character, and byte count",
		UsageText:   "wc [-clmw] [file ...]",
		Description: description,
		Action: func(c *cli.Context) error {
			fmt.Println("wc")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
	}
}
