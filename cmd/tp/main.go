package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/cwd-k2/gvfs"
	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	logger = log.New(os.Stderr, "[tp:github] ", log.LstdFlags|log.Ltime)
	cmd    = &cobra.Command{
		Use:  "tp <owner>/<repo> <directory-name>",
		Args: cobra.ExactArgs(2),
		RunE: FromGitHub,
	}
)

func FromGitHub(_ *cobra.Command, args []string) error {
	tmp, err := ioutil.TempDir("", "templative-")
	if err != nil {
		return errors.WithStack(err)
	}

	git.PlainClone(tmp, false, &git.CloneOptions{
		URL:      "https://github.com/" + args[0],
		Progress: os.Stdout,
	})

	dstdir, err := filepath.Abs(args[1])
	if err != nil {
		return errors.WithStack(err)
	}

	src, err := gvfs.Traverse(tmp, regexp.MustCompile(`\.git$`))
	if err != nil {
		return errors.WithStack(err)
	}

	for _, content := range src.Contents {
		if err := content.Commit(dstdir); err != nil {
			logger.Printf("%+v\n", errors.WithStack(err))
		}
	}

	return nil
}

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Printf("%+v\n", err)
		os.Exit(1)
	}
}
