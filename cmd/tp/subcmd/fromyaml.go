package subcmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cwd-k2/gvfs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	logger   = log.New(os.Stderr, "[tp:yaml] ", log.LstdFlags|log.Ltime)
	FromYaml = &cobra.Command{
		Use:  "yaml <path-to-yaml> <directory-name>",
		Args: cobra.ExactArgs(2),
		RunE: fromyaml,
	}
)

const (
	EmptyDirectory = "(directory)"
	EmptyFile      = "(file)"
)

func fromyaml(_ *cobra.Command, args []string) error {
	yamlpath, err := filepath.Abs(args[0])
	if err != nil {
		return errors.WithStack(err)
	}

	abstarget, err := filepath.Abs(args[1])
	if err != nil {
		return errors.WithStack(err)
	}

	yamlbyte, err := ioutil.ReadFile(yamlpath)
	if err != nil {
		return errors.WithStack(err)
	}

	var structure map[string]string
	if err := yaml.Unmarshal(yamlbyte, &structure); err != nil {
		return errors.WithStack(err)
	}

	var (
		parent, basename = filepath.Split(abstarget)
		directory        = gvfs.NewDirectory(basename)
	)

	for path, content := range structure {
		switch content {
		case EmptyDirectory:
			if _, err := directory.CreateDirectory(gvfs.NewPath(path)); err != nil {
				logger.Printf("%+v\n", errors.WithStack(err))
			}
		case EmptyFile:
			if _, err := directory.CreateFile(gvfs.NewPath(path)); err != nil {
				logger.Printf("%+v\n", errors.WithStack(err))
			}
		default:
			file, err := directory.CreateFile(gvfs.NewPath(path))
			if err != nil {
				logger.Printf("%+v\n", errors.WithStack(err))
				continue
			}
			if _, err := strings.NewReader(content).WriteTo(file); err != nil {
				logger.Printf("%+v\n", errors.WithStack(err))
			}
		}
	}

	if err := directory.Commit(parent); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
