// Package checks provides functions check passed into CLI arguments
package checks

import "regexp"

// CheckRepositoryURL checks for correctness repository url.
func CheckRepositoryURL(url string) error {
	match, _ := regexp.MatchString(`^https://github.com/.*$`, url)
	if !match {
		return InvalidInputError{"repository URL"}
	}

	return nil
}

// CheckExtension checks for correctness files extension.
func CheckExtension(ext string) error {
	match, _ := regexp.MatchString(`^\.\S*$`, ext)
	if !match {
		return InvalidInputError{"file extension"}
	}

	return nil
}

// CheckFolderPrefix checks for correctness folder markdown prefix.
func CheckFolderPrefix(folderPrefix string) error {
	if len(folderPrefix) == 0 {
		return InvalidInputError{"folder-prefix"}
	}

	return nil
}

// CheckFilePrefix checks for correctness file markdown prefix.
func CheckFilePrefix(folderPrefix string) error {
	if len(folderPrefix) == 0 {
		return InvalidInputError{"file-prefix"}
	}

	return nil
}

// CheckInputData checks for correctness all passed arguments.
func CheckInputData(url, ext, folderPrefix, filePrefix string) error {
	err := CheckRepositoryURL(url)
	if err != nil {
		return err
	}

	err = CheckExtension(ext)
	if err != nil {
		return err
	}

	err = CheckFolderPrefix(ext)
	if err != nil {
		return err
	}

	err = CheckFilePrefix(ext)
	if err != nil {
		return err
	}

	return nil
}
