//go:build ignore

// example on how to change help args' position in help message
//
// note that the order of your argument is the same as the order when you add them,
// then you can decide the order of your arguments' display
package main

import (
	"fmt"

	"github.com/hellflame/argparse"
)

func main() {
	parse := argparse.NewParser("", "", &Argparse.ParserConfig{DisableHelp: true})
	parse.Strings("f", "first", &Argparse.Option{Help: "this is first"})
	parse.String("s", "second", &Argparse.Option{Help: "this is second"})
	parse.Flag("h", "help", &Argparse.Option{
		Help: "show this help at bottom",
		Action: func(a []string) error {
			fmt.Println("============")
			parse.PrintHelp()
			return nil
		},
	})
	parse.Parse(nil)
}
