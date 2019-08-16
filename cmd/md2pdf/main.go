package main

import (
	"fmt"

	"github.com/tcd/md2pdf/internal/ghfm"
	"github.com/tcd/md2pdf/md2pdf"
)

func main() {
	runParse()
}

const (
	pdfOutDir    = "../../out/pdf/"
	htmlOutDir   = "../../out/html/"
	readme       = "../../readme.md"
	mdCheatsheet = "../../static/markdown-cheatsheet.md"
)

func runParse() {
	ghfm.GitHubPDF(pdfOutDir + "table3.pdf")
	// parse.GenAST()

	err := md2pdf.Md2PDF(mdCheatsheet, pdfOutDir+"github3.pdf")
	if err != nil {
		fmt.Println(err)
	}
}
