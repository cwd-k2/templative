package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/cwd-k2/gvfs"
	"github.com/go-git/go-git/v5"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	templatepath, err := ioutil.TempDir("", "templative-")
	if err != nil {
		panic(err)
	}

	git.PlainClone(templatepath, false, &git.CloneOptions{
		URL:      "https://github.com/" + os.Args[1],
		Progress: os.Stdout,
	})

	dstdir, err := filepath.Abs(os.Args[2])
	if err != nil {
		panic(err)
	}

	src, err := gvfs.Traverse(templatepath, regexp.MustCompile(`\.git$`))
	if err != nil {
		panic(err)
	}

	for _, content := range src.Contents {
		if err := content.Commit(dstdir); err != nil {
			println(err.Error())
		}
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: tp <owner>/<repo> <directory-name>`)
	os.Exit(1)
}
