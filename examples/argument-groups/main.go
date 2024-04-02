//go:build ignore

// this is show case for creating argument groups
package main

import (
	"fmt"
	"os"

	"github.com/hellflame/argparse"
)

func main() {
	p := argparse.NewParser("", "this is a show case about groups", &Argparse.ParserConfig{DisableHelp: true, EpiLog: "try more"})
	p.Flag("n", "normal", nil)
	p.Float("f", "float", &Argparse.Option{Positional: true})

	p.String("a", "aa", &Argparse.Option{Group: "As", Help: "normal a"})
	p.String("aaa", "", &Argparse.Option{Group: "As", Help: "triple a", Positional: true})

	p.Int("b", "bb", &Argparse.Option{Group: "Bs", Help: "normal b"})
	p.Ints("", "bbb", &Argparse.Option{Group: "Bs", Help: "triple b"})

	help := p.Flag("h", "", &Argparse.Option{Group: "General", Help: "show help info"})
	if e := p.Parse(nil); e != nil {
		switch e.(type) {
		case argparse.BreakAfterHelp:
			os.Exit(1)
		default:
			fmt.Println(e.Error())
			return
		}
	}
	if *help {
		p.PrintHelp()
		return
	}
}
