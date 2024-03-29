package md2pdf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tcd/md2pdf/internal/parse"
	"github.com/tcd/md2pdf/internal/renderer"
	rdr "github.com/tcd/md2pdf/internal/renderers/github"
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
	rbs := parse.Parse(bytes)
	var r rdr.Renderer
	err = renderer.RenderToFile(r, rbs, newFile)
	if err != nil {
		return "", err
	}
	return newFile, nil
}

// MdURLToPdfFile converts a markdown file at a given URL to a pdf file.
// The path to the new PDF file is returned along with any encountered errors.
func MdURLToPdfFile(url, outPath string) (string, error) {
	bytes, err := bytesFromURL(url)
	if err != nil {
		return outPath, err
	}
	bytes = mdBytes2htmlbytes(bytes)
	rbs := parse.Parse(bytes)
	var r rdr.Renderer
	err = renderer.RenderToFile(r, rbs, outPath)
	if err != nil {
		return outPath, err
	}
	return outPath, nil
}

// MdBytesToPdfFile can be used to render markdown input from stdin.
// The path to the new PDF file is returned along with any encountered errors.
func MdBytesToPdfFile(mdBytes []byte, outPath string) (string, error) {
	if outPath == "" {
		outPath = "no-name.pdf"
	}
	newFile := absPath(outPath)
	ensureDir(newFile)
	htmlBytes := mdBytes2htmlbytes(mdBytes)
	rbs := parse.Parse(htmlBytes)
	var r rdr.Renderer
	err := renderer.RenderToFile(r, rbs, newFile)
	if err != nil {
		return "", err
	}
	return newFile, nil
}

// Debug outputs not only a PDF, but also HTML and JSON output for debugging.
func Debug(path, debugDir string) error {
	if debugDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		debugDir = filepath.Join(cwd, "debug")
	}
	oldFile := absPath(path)
	baseName := replaceExtension(oldFile, "")
	outDir := filepath.Join(debugDir, baseName)
	makeDir(outDir)
	mdOut := filepath.Join(outDir, baseName) + ".md"
	htmlOut := filepath.Join(outDir, baseName) + ".html"
	jsonOut := filepath.Join(outDir, baseName) + ".json"
	pdfOut := filepath.Join(outDir, baseName) + ".pdf"

	// Copy the markdown file.
	err := copyFile(path, mdOut)
	if err != nil {
		return err
	}

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
	rbs := parse.Parse(htmlBytes)

	// Write the JSON output.
	bytes, err := json.MarshalIndent(rbs, "", "  ")
	err = ioutil.WriteFile(jsonOut, bytes, os.FileMode(0644))
	if err != nil {
		return err
	}

	// Write the PDF output.
	var r rdr.Renderer
	err = renderer.RenderToFile(r, rbs, pdfOut)
	if err != nil {
		return err
	}

	log.Println("Output in:", outDir)
	return nil
}
