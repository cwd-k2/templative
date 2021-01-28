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
	templatepath := filepath.Join(os.TempDir(), "templattive", uuidobj.String())

	git.PlainClone(templatepath, false, &git.CloneOptions{
		URL:      "https://github.com/" + os.Args[1],
		Progress: os.Stdout,
	})

	dstDir, err := filepath.Abs(os.Args[2])
	if err != nil {
		panic(err)
	}

	src := gvfs.NewRoot(templatepath)
	dst := gvfs.NewRoot(dstDir)

	dir, err := src.ToItem(regexp.MustCompile(`\.git`))
	if err != nil {
		panic(err)
	}

	for _, content := range dir.Contents {
		if err := dst.WriteItem(content); err != nil {
			println(err)
		}
	}

}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: tp <owner>/<repo> <directory-name>`)

	os.Exit(1)
}
