package md2pdf

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// copyFile copies the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// Return the contents of a file as a byte slice.
func bytesFromFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}

func bytesFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
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

// Ensure a directory exists, create it if necessary.
func ensureDir(path string) {
	folder := filepath.Dir(absPath(path))
	fmt.Println(folder)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.ModePerm)
	}
}

// Create a directory if it doesn't exist.
func makeDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
