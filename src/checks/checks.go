package checks

import "regexp"

func CheckRepositoryURL(url string) error {
	match, _ := regexp.MatchString(`^https://github.com/.*$`, url)
	if !match {
		return InvalidInputError{"repository URL"}
	}

	return nil
}

func CheckExtension(ext string) error {
	match, _ := regexp.MatchString(`^\.\S*$`, ext)
	if !match {
		return InvalidInputError{"file extension"}
	}

	return nil
}

func CheckInputData(url, ext string) error {
	err := CheckRepositoryURL(url)
	if err != nil {
		return err
	}

	err = CheckExtension(ext)
	if err != nil {
		return err
	}

	return nil
}
