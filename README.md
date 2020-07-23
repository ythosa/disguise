# disguise

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/db828996dabb4d2a9e00c1fb3263bcd4)](https://app.codacy.com/manual/Ythosa/disguise?utm_source=github.com&utm_medium=referral&utm_content=Ythosa/disguise&utm_campaign=Badge_Grade_Dashboard)

Disguise is CLI tool for generation markdown  with list of github repository directories and files. Can be used for creation repositories issues about the process of documenting the code.

Usage example:
* ./disguise -help
* ./disguise [options] -url "<repository_name>" -ext "<files_extension>".

Options could be: 
* -url "<github_repo_url>" - identify repositories should have documentation
* -ext "<file_extension>" - identify files should have documentation
* -ignore "<some_dir_name_in_repo>" - identify dirs shouldn't have documentation
* -help - returns some CLI help and info
