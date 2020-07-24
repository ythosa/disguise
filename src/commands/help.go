// Package commands provides functions for CLI commands
package commands

import "fmt"

// GetHelp returns help for help CLI command
func GetHelp() {
	fmt.Println(`			~ Welcome to the disguise club buddy. ~
Disguise is CLI tool for generation markdown  with list of github repository directories and files.
Can be used for creation repositories issues about the process of documenting the code.
Usage example:
	./disguise [options] -url \"<repository_name>\" -ext \"<files_extension>\".
Options could be: \n\t" +
	-url \"<github_repo_url>\" - identify repositories should have documentation
	-ext \"<file_extension>\" - identify files should have documentation
	-ignore \"<some_dir_name_in_repo>\" - identify dirs shouldn't have documentation
	-help - returns some CLI help and info
Author: Ythosa <vasus714@yandex.ru> https://github.com/Ythosa`)
}
