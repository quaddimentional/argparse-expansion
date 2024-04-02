//go:build ignore

// this is show case for using color in argparse
//
// demo for ColorSchema
package main

import (
	"fmt"

	"github.com/hellflame/argparse"
)

func main() {
	parser := argparse.NewParser("basic", "this is a basic program", &Argparse.ParserConfig{
		WithColor: true,
		WithHint:  true,

		EpiLog: "more info please visit https://github.com/hellflame/argparse",
	})
	sub := parser.AddCommand("run", "run your program", nil)
	parser.AddCommand("test", "test for your program", nil)

	sub.Flag("d", "dir", &Argparse.Option{Help: "give me a directory"})
	parser.String("n", "name", &Argparse.Option{Default: "flame", Help: "your name"})
	parser.Ints("t", "times", &Argparse.Option{HintInfo: "run times", Group: "base", Help: "how many times"})
	parser.Float("s", "size", &Argparse.Option{Help: "give me a size", Group: "base", Required: true})
	parser.String("u", "url", &Argparse.Option{Positional: true, Help: "target url"})
	parser.String("l", "", nil)
	if e := parser.Parse(nil); e != nil {
		switch e {
		case argparse.BreakAfterHelpError:
			return
		default:
			fmt.Println(e.Error())
		}
		return
	}
}
