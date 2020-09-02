你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
<br>

<div align="center">
<h1>disguise</h1>

[![CI/CD](https://github.com/Ythosa/disguise/workflows/Go/badge.svg?branch=master)](https://github.com/Ythosa/disguise/actions)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/db828996dabb4d2a9e00c1fb3263bcd4)](https://app.codacy.com/manual/Ythosa/disguise?utm_source=github.com&utm_medium=referral&utm_content=Ythosa/disguise&utm_campaign=Badge_Grade_Dashboard)
[![CodeFactor](https://www.codefactor.io/repository/github/ythosa/disguise/badge)](https://www.codefactor.io/repository/github/ythosa/disguise)
</div>

<br>

<img src="https://camo.githubusercontent.com/98ed65187a84ecf897273d9fa18118ce45845057/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67" alt="gopher :3" height="200px"/>

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
_Q_: How can I help to develop this project?  
_A_: You can put a :star: :3

<br>

<div align="center">
  Copyright 2020 Ythosa
</div>
