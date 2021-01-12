package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cwd-k2/templative/pkg/constants"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	TEMPLATIVE_DIR := constants.TemplativeDir()

	if _, err := os.Stat(TEMPLATIVE_DIR); os.IsNotExist(err) {
		os.Mkdir(TEMPLATIVE_DIR, 0755)
	}

	templatePath := filepath.Join(TEMPLATIVE_DIR, os.Args[1]) + string(filepath.Separator)

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		repo := "https://github.com/" + os.Args[1]

		cmd := exec.Command("git", "clone", repo, templatePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
	}

	d, _ := filepath.Abs(os.Args[2])

	cmd := exec.Command("rsync", "-av", "--cvs-exclude", templatePath, d)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: tp <owner>/<repo> <directory-name>`)

	os.Exit(1)
}
