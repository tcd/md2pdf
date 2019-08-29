package md2pdf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tcd/md2pdf/internal/parse"
)

// MdFileToPdfFile converts a markdown file to a PDF file.
// The path to the new PDF file is returned along with any encountered errors.
func MdFileToPdfFile(inPath, outPath string) (string, error) {
	var oldFile, newFile string
	if outPath == "" {
		// Generate a name if we aren't given one.
		oldFile, newFile = parsePaths(inPath, ".pdf")
	} else {
		oldFile = absPath(inPath)
		newFile = absPath(outPath)
		ensureDir(newFile)
	}

	bytes, err := mdFile2htmlBytes(oldFile)
	if err != nil {
		return "", err
	}
	elements := parse.Parse(bytes)
	err = elements.RenderToFile(newFile)
	if err != nil {
		return "", err
	}
	return newFile, nil
}

// MdFileToUnnamedPdf converts a markdown file to a PDF file.
// The path to the new PDF file is returned along with any encountered errors.
func MdFileToUnnamedPdf(path string) (string, error) {
	oldFile, newFile := parsePaths(path, ".pdf")
	bytes, err := mdFile2htmlBytes(oldFile)
	if err != nil {
		return newFile, err
	}
	elements := parse.Parse(bytes)
	err = elements.RenderToFile(newFile)
	if err != nil {
		return newFile, err
	}
	return newFile, err
}

// Debug outputs not only a PDF, but also HTML and JSON output for debugging.
func Debug(path, debugDir string) error {
	if debugDir == "" {
		debugDir = "/Users/clay/go/src/github.com/tcd/md2pdf/out/debug"
	}
	oldFile := absPath(path)
	baseName := replaceExtension(oldFile, "")
	outDir := filepath.Join(debugDir, baseName)
	makeDir(outDir)
	htmlOut := filepath.Join(outDir, baseName) + ".html"
	jsonOut := filepath.Join(outDir, baseName) + ".json"
	pdfOut := filepath.Join(outDir, baseName) + ".pdf"

	// Parse the markdown file.
	htmlBytes, err := mdFile2htmlBytes(oldFile)
	if err != nil {
		return err
	}

	// Write the html output.
	err = ioutil.WriteFile(htmlOut, htmlBytes, os.FileMode(0644))
	if err != nil {
		return err
	}

	// Parse the html.
	elements := parse.Parse(htmlBytes)

	// Write the JSON output.
	bytes, err := json.Marshal(elements)
	err = ioutil.WriteFile(jsonOut, bytes, os.FileMode(0644))
	if err != nil {
		return err
	}

	// Write the PDF output.
	err = elements.RenderToFile(pdfOut)
	if err != nil {
		return err
	}

	log.Println("Output in:", outDir)
	return nil
}
