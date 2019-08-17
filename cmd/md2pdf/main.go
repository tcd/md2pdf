package main

import (
	"fmt"

	"github.com/tcd/md2pdf/internal/md2pdf"
)

func main() {
	runParse()
}

const (
	pdfOutDir    = "../../out/pdf/"
	htmlOutDir   = "../../out/html/"
	readme       = "../../readme.md"
	mdCheatsheet = "../../testdata/markdown/markdown-cheatsheet.md"
)

func runParse() {
	// md2pdf.GitHubPDF(pdfOutDir + ".pdf")
	// parse.GenAST()

	err := md2pdf.Md2PDF(mdCheatsheet, pdfOutDir+"newest.pdf")
	if err != nil {
		fmt.Println(err)
	}
}
