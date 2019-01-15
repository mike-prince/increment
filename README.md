<h1 align="center">increment</h1>

<p align="center">CLI tool for renaming files by appending/prepending (or replacing) the filename with incremental count.</p>

### Installation

Via *Go*:

`go get github.com/mike-prince/increment`

### Usage

`increment [options] <glob pattern>`

##### Options

- `-a  append count to existing file name`
- `-p  prepend count to existing file name`
- `-c  number to begin count`
- `-t  output result without making changes`
- `-o  output result`
- `-h  output this help and exit`
- `-V  output the version and exit`

There are no safeguards in place. Always make a backup before running as files will be renamed if the user has permission. :bomb:

### Examples

Given files in current directory *foo.txt bar.txt baz*

```
Command:  increment *
Result:   1.txt 2.txt 3
```
```
Command:  increment -p -c=6 "*.txt"
Result:   6foo.txt 7bar.txt 8baz
```
```
Command:  increment -a "b*"
Result:   bar1.txt baz2
```
```
Command:  increment -a -p *
Result:   1.txt 2.txt 3
```
