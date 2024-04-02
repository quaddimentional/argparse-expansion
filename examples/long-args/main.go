//go:build ignore

// example on how to deal with very long args (if it's very necessary)
//
// MaxHeaderLength is recommended to be around 20 ~ 30
package main

import "github.com/hellflame/argparse"

func main() {
	parser := argparse.NewParser("long-args", "", &Argparse.ParserConfig{
		MaxHeaderLength: 20, // if argument header length exceeds 20 characters, argument help message will display on new line with 20 space indent
	})
	parser.String("s", "short", &Argparse.Option{Help: "this is a short args"})
	parser.String("m", "medium-size", &Argparse.Option{Help: "this is a medium size args"})
	parser.String("l", "this-is-a-very-long-args", &Argparse.Option{Help: "this is a very long args"})
	parser.Parse(nil)
}
