# disguise

[![CodeFactor](https://www.codefactor.io/repository/github/ythosa/disguise/badge)](https://www.codefactor.io/repository/github/ythosa/disguise)

Disguise is CLI tool for generation markdown  with list of github repository directories and files. Can be used for creation repositories issues about the process of documenting the code.

Usage example:
* ./disguise -help
* ./disguise [options] -url "<repository_name>" -ext "<files_extension>".

Options could be: 
* -url "<github_repo_url>" - identify repositories should have documentation
* -ext "<file_extension>" - identify files should have documentation
* -ignore "<some_dir_name_in_repo>" - identify dirs shouldn't have documentation
* -help - returns some CLI help and info
