package constants

import (
	"os"
	"path/filepath"
)

// TEMPLATIVE_DIR 環境変数を頑張って読む
func TemplativeDir() string {
	TEMPLATIVE_DIR := os.Getenv("TEMPLATIVE_DIR")

	if len(TEMPLATIVE_DIR) > 0 {
		return TEMPLATIVE_DIR
	}

	XDG_DATA_HOME := os.Getenv("XDG_DATA_HOME")

	if len(XDG_DATA_HOME) > 0 {
		return filepath.Join(XDG_DATA_HOME, "templative")
	}

	HOME := os.Getenv("HOME")

	return filepath.Join(HOME, ".local", "share", "templative")
}
