package main

import (
	"github.com/tcd/md2pdf/internal/md2pdf"
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
	md2pdf.GitHubPDF(pdfOutDir + "list2.pdf")
	// parse.GenAST()

	// err := md2pdf.Md2PDF(mdCheatsheet, pdfOutDir+"newest.pdf")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
