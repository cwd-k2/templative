package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/cwd-k2/gvfs"
	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	uuidobj, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	templatepath := filepath.Join(os.TempDir(), "templative", uuidobj.String())

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

	dst := gvfs.NewDirectory(filepath.Base(dstdir))
	dst.Contents = src.Contents

	if err := dst.Commit(filepath.Dir(dstdir)); err != nil {
		println(err.Error())
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: tp <owner>/<repo> <directory-name>`)
	os.Exit(1)
}
