package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var version = "0.1.0"

const usage = `
  Usage:
    increment [options] <glob pattern>

  Options:
    -a	append count to existing file name
    -p	prepend count to existing file name
    -c	number to begin count
    -t	output result without making changes
    -o	output result
    -h	output this help and exit
    -V	output the version and exit

  Examples:
    Given files in current directory foo.txt bar.txt baz 
    		
    Command:	increment *
    Result:	1.txt 2.txt 3
    
    Command:	increment -p -c=6 "*.txt"
    Result:	6foo.txt 7bar.txt 8baz
    
    Command:	increment -a "b*"
    Result:	bar1.txt baz2

    Command:	increment -a -p *
    Result:	1.txt 2.txt 3

  Warning:
    There are no safeguards in place. Always make a backup before
    running as files will be renamed if the user has permission.
`

type File struct {
	name    string
	ext     string
	oldName string
	newName string
}

func main() {

	argA := flag.Bool("a", false, "append bool")
	argP := flag.Bool("p", false, "prepend bool")
	argC := flag.Int("c", 1, "count integer")
	argT := flag.Bool("t", false, "test bool")
	argV := flag.Bool("v", false, "verbose bool")
	argH := flag.Bool("h", false, "help bool")
	argVersion := flag.Bool("V", false, "version string")

	flag.Parse()
	files := flag.Args()

	switch {
	case *argVersion == true:
		fmt.Println(version)
		return
	case *argH == true || len(files) == 0:
		fmt.Println(usage)
		return
	}

	position := ""

	switch {
	case *argP == true && *argA == false:
		position = "p"
	case *argP == false && *argA == true:
		position = "a"
	}

	count := *argC

	for _, f := range files {

		e := strings.Split(f, ".")
		file := File{name: e[0]}
		file.oldName = f

		switch position {
		case "p":
			file.newName = fmt.Sprintf("%d%s", count, file.name)
		case "a":
			file.newName = fmt.Sprintf("%s%d", file.name, count)
		default:
			file.newName = fmt.Sprintf("%d", count)
		}

		if len(e) > 1 {
			file.ext = "." + e[1]
		}

		file.newName = file.newName + file.ext

		switch {
		case *argT == true:
			fmt.Printf("%s > %s\n", file.oldName, file.newName)
		case *argV == true:
			fmt.Printf("%s > %s\n", file.oldName, file.newName)
			os.Rename(file.oldName, file.newName)
		default:
			os.Rename(file.oldName, file.newName)
		}

		count++
	}
}
