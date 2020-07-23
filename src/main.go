package main

import (
	"flag"
	"github.com/ythosa/disguise/src/commands"
)

var help = flag.Bool("help", false, "Returns help with CLI.")
var url = flag.String("url", "", "Which repository should have documentation.")
var extension = flag.String("ext", "", "Which files should have documentation")
var toIgnore = flag.String("ignore", "", "Which dirs shouldn't have documentation.")

/* Example of starting program
	<./cli_path> [options] -url "<repository_name>" -ext "<files_extension>"
	./disguise -ignore "Platform.Setters.Tests" -url https://github.com/linksplatform/Setters/ --ext ".cs"
	./disguise -help
*/
func main() {
	flag.Parse()

	if *help {
		commands.GetHelp()
		return
	}

	if (len(*url) != 0) || (len(*extension) != 0) {
		commands.GetMarkdown(*url, *extension, *toIgnore)
		return
	}

	commands.GetHelp()
}
