// Package checks provides functions check passed into CLI arguments
package checks

import "regexp"

// checkRepositoryURL checks for correctness repository url
func checkRepositoryURL(url string) error {
	match, _ := regexp.MatchString(`^https://github.com/.*$`, url)
	if !match {
		return invalidInputError{"repository URL"}
	}

	return nil
}

// checkExtension checks for correctness files extension
func checkExtension(ext string) error {
	match, _ := regexp.MatchString(`^\.\S*$`, ext)
	if !match {
		return invalidInputError{"file extension"}
	}

	return nil
}

// CheckInputData checks for correctness all passed arguments
func CheckInputData(url, ext string) error {
	err := checkRepositoryURL(url)
	if err != nil {
		return err
	}

	err = checkExtension(ext)
	if err != nil {
		return err
	}

	return nil
}
