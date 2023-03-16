package funcs

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Aeswolf/cdx/cd"
)

func walk(root string, cdx *cd.CDX) (p string) {
	src, dest := cdx.GetSrc(), cdx.GetDest()

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && isHidden(path) {
			return filepath.SkipDir
		}

		if containsSD(path, src, dest) && hasSameBase(path, dest) {
			p = path
			return nil
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking dir from the root %s ===> %s", root, err.Error())
	}

	return
}

// function to check for hidden files
func isHidden(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

// function to get if a path contains both the src and dest of cdx
func containsSD(path, src, dest string) bool {
	return strings.Contains(path, src) && strings.Contains(path, dest)
}

// function to ensure base paths are the same
func hasSameBase(path, dest string) bool {
	return filepath.Base(path) == filepath.Base(dest)
}

// function to the get the required path
func GetPath(cdx *cd.CDX) string {
	src := cdx.GetSrc()

	root, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error attaining the path to the current working directory ===> %s", err.Error())
	}

	if src != "." {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			log.Fatalf("Error accessing home directory path ===> %s", err.Error())
		}

		root = homeDir
	}

	return walk(root, cdx)
}
