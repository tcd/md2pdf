package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// return a nice Ascii title for "md2pdf"
var titleString = strings.Join(titleLines, "\n")
var titleLines = []string{
	"               _  _____           _  __",
	"              | |/ __  \\         | |/ _|",
	" _ __ ___   __| |`' / /'_ __   __| | |_",
	"| '_ ` _ \\ / _` |  / / | '_ \\ / _` |  _|",
	"| | | | | | (_| |./ /__| |_) | (_| | |",
	"|_| |_| |_|\\__,_|\\_____/ .__/ \\__,_|_|",
	"                       | |",
	"                       |_|",
}

// Ensure a directory exists, create it if necessary.
func ensureDir(path string) {
	folder := filepath.Dir(absPath(path))
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.ModePerm)
	}
}

// parsePaths returns an absolute path to the given existing file,
// and the full path of the new file to create.
func parsePaths(inFile, ext string) (string, string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	oldFile := absPath(inFile)
	newFile := filepath.Join(cwd, replaceExtension(oldFile, ext))
	return oldFile, newFile
}

// return an absolute path, or the given path if it's already absolute.
func absPath(path string) string {
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		return absPath
	}
	return path
}

// return a new filename with a replaced extension.
func replaceExtension(file, ext string) string {
	baseName := filepath.Base(file)
	extension := filepath.Ext(baseName)
	newFileName := strings.Replace(baseName, extension, ext, -1)
	return newFileName
}
