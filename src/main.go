package main

import (
	"flag"
	"fmt"

	"github.com/ythosa/disguise/src/checks"
	"github.com/ythosa/disguise/src/commands"
)

/*
Example of starting program:L
	<./cli_path> [options] -url "<repository_name>" -ext "<files_extension>"
	./disguise -ignore "Platform.Setters.Tests" -url https://github.com/linksplatform/Setters/ --ext ".cs"
	./disguise -help
*/
func main() {
	var help = flag.Bool("help", false, "Returns help with CLI.")
	var url = flag.String("url", "", "Which repository should have documentation.")
	var extension = flag.String("ext", "", "Which files should have documentation")
	var toIgnore = flag.String("ignore", "", "Which dirs shouldn't have documentation.")
	var folderPrefix = flag.String("folder-prefix", "#####", "Folder prefix in markdown.")
	var filePrefix = flag.String("file-prefix", "- [ ]", "File prefix in markdown.")

	flag.Parse()

	if *help {
		commands.GetHelp()
		return
	}

	if (len(*url) != 0) || (len(*extension) != 0) {
		err := checks.CheckInputData(*url, *extension, *folderPrefix, *filePrefix)
		if err != nil {
			fmt.Printf("error. %s. \n"+
				"use -help flag to get using template.\n", err.Error())
			return
		}
		commands.GetMarkdown(*url, *extension, *toIgnore,
			commands.MarkDownConfig{
				Files: *filePrefix,
				Dirs:  *folderPrefix,
			})
		return
	}

	commands.GetHelp()
}
