//go:build ignore

// this show case is for Argument Groups
//
// group arguments in different group in help message
//
// but if there is too many argument, it's better to change ParserConfig.Usage to a simpler string
//
// # Argument Group is for user to better understand your program group by group
//
// you should also checkout SubCommand for reference, they are quit like each other
package main

import (
	"fmt"
	"strings"

	"github.com/hellflame/argparse"
)

func main() {
	parser := argparse.NewParser("yt-download", "youtube downloader, download anything on youtube", &Argparse.ParserConfig{Usage: "yt-download [OPTIONS] URL [URL...]", EpiLog: "bug report to: xxx@xxx.xx"})
	showVersion := parser.Flag("", "version", &Argparse.Option{Help: "Print program version and exit", Group: "GeneralOptions"})
	doUpdate := parser.Flag("U", "update", &Argparse.Option{Help: "Update this program to latest version. Make sure that you have sufficient permissions (run with sudo if needed)", Group: "GeneralOptions"})
	listExtractors := parser.Flag("", "list-extractors", &Argparse.Option{Help: "List all supported extractors", Group: "GeneralOptions"})
	search := parser.String("", "default-search", &Argparse.Option{Help: "Use this prefix for unqualified URLs. For example \"gvsearch2:\" downloads two videos from google videos for youtube-dl \"large apple\". Use the value \"auto\" to let youtube-dl", Meta: "PREFIX"})

	playListStart := parser.Int("", "playlist-start", &Argparse.Option{Help: "Playlist video to start at (default is 1)", Meta: "NUMBER", Group: "play-list options", Default: "1"})

	urls := parser.Strings("", "url", &Argparse.Option{Help: "youtube links, like 'https://www.youtube.com/watch?v=xxxxxxxx'", Positional: true, Required: true, Validate: func(arg string) error {
		if !strings.HasPrefix(arg, "https://") {
			return fmt.Errorf("url should be start with 'https://'")
		}
		return nil
	}})

	if e := parser.Parse(nil); e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(*showVersion, *doUpdate, *listExtractors, *search, *playListStart, len(*urls))
}
