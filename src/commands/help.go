// Package commands provides functions for CLI commands
package commands

import "fmt"

// GetHelp returns help for help CLI command.
func GetHelp() {
	fmt.Println(`& Disguise ~
Disguise is CLI tool for generation markdown with list of github repository directories and files.
Can be used for creation repositories issues about the process of documenting the code.
Usage example:
	./disguise [options] -url "<repository_name>" -ext "<files_extension>".
Options could be:
	-url "<github_repo_url>" - identify repositories should have documentation
	-ext "<file_extension>" - identify files should have documentation
	-ignore "<some_dir_name_in_repo>" - identify dirs shouldn't have documentation
	-file-prefix "<some_markdown_syntax>" - specifies the prefix of files in the output markdown (default: - [ ])
	-folder-prefix "<some_markdown_syntax>" - specifies the prefix of folders in the output markdown (default: #####)
	
	-help - returns some CLI help and info
Author: Ythosa <vasus714@yandex.ru> https://github.com/Ythosa`)
}
