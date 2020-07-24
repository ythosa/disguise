// Package commands provides functions for CLI commands
package commands

import "fmt"

// GetHelp returns help for help CLI command
func GetHelp() {
	fmt.Println("\t~ Welcome to the disguise club buddy. ~\n" +
		"Disguise is CLI tool for generation markdown  with list of github repository directories and files." +
		"Can be used for creation repositories issues about the process of documenting the code.\n\n" +
		"Usage example:\n\t" +
		"./disguise [options] -url \"<repository_name>\" -ext \"<files_extension>\".\n" +
		"Options could be: \n\t" +
		"-url \"<github_repo_url>\" - identify repositories should have documentation\n\t" +
		"-ext \"<file_extension>\" - identify files should have documentation\n\t" +
		"-ignore \"<some_dir_name_in_repo>\" - identify dirs shouldn't have documentation\n\t" +
		"-help - returns some CLI help and info\n\n" +
		"Author: Ythosa <vasus714@yandex.ru> https://github.com/Ythosa")
}
