package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/cwd-k2/gvfs"
	"github.com/cwd-k2/templative/pkg/constants"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	TEMPLATIVE_DIR := constants.TemplativeDir()

	if _, err := os.Stat(TEMPLATIVE_DIR); os.IsNotExist(err) {
		os.Mkdir(TEMPLATIVE_DIR, os.ModePerm)
	}

	templatePath := filepath.Join(TEMPLATIVE_DIR, os.Args[1]) + string(filepath.Separator)

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		repo := "https://github.com/" + os.Args[1]

		cmd := exec.Command("git", "clone", repo, templatePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
	}

	dstDir, err := filepath.Abs(os.Args[2])
	if err != nil {
		panic(err)
	}

	src := gvfs.NewRoot(templatePath)
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
