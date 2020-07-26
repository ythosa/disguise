<br>

<div align="center">
<h1>disguise</h1>

[![CI/CD](https://github.com/Ythosa/disguise/workflows/Go/badge.svg?branch=master)](https://github.com/Ythosa/disguise/actions)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/db828996dabb4d2a9e00c1fb3263bcd4)](https://app.codacy.com/manual/Ythosa/disguise?utm_source=github.com&utm_medium=referral&utm_content=Ythosa/disguise&utm_campaign=Badge_Grade_Dashboard)
[![CodeFactor](https://www.codefactor.io/repository/github/ythosa/disguise/badge)](https://www.codefactor.io/repository/github/ythosa/disguise)
</div>

<br>

## Install
*   From sources:
    *   Clone this repo: `git clone https://github.com/Ythosa/disguise`;
    *   Build project by writing in terminal: `make`.
*   From github tag:
    *   Download the [latest](https://github.com/Ythosa/disguise/releases) binaries of this project.
      
## Description
Disguise is CLI tool for generation markdown with list of github repository directories and files. 
Can be used for creation repositories issues about the process of documenting the code.

CLI options could be:
*   -url "<github_repo_url>" - identify repositories should have documentation
*   -ext "<file_extension>" - identify files should have documentation
*   -ignore "<some_dir_name_in_repo>" - identify dirs shouldn't have documentation
*   -file-prefix "<some_markdown_syntax>" - specifies the prefix of files in the output markdown (default: - \[ \])
*   -folder-prefix "<some_markdown_syntax>" - specifies the prefix of folders in the output markdown (default: #####)
*   -help - returns CLI help and info

## Usage example
Let's create an issue about the process of documenting this project, but I don't want to document the _checks_ folder.
We just need to write this line in the terminal.
```shell script
./disguise -url "https://github.com/ythosa/disguise" -ext ".go" -ignore "src/checks" -folder-prefix "*" -file-prefix "    *"
```
Output example:
*   [src](https://github.com/Ythosa/disguise/tree/master/src)
    *   [main](https://github.com/Ythosa/disguise/blob/master/src/main.go)

*   [src/commands](https://github.com/Ythosa/disguise/tree/master/src/commands)
    *   [help](https://github.com/Ythosa/disguise/blob/master/src/commands/help.go)
    *   [markdown](https://github.com/Ythosa/disguise/blob/master/src/commands/markdown.go)
    *   [markdown_test](https://github.com/Ythosa/disguise/blob/master/src/commands/markdown_test.go)

## FAQ
*Q*: How can I help to develop this project?  
*A*: You can put a :star: :3

<br>

<div align="center">
  Copyright 2020 Ythosa
</div>
