package main

import (
	"fmt"

	"github.com/tcd/md2pdf/internal/md2pdf"
)

const (
	pdfOutDir    = "../../out/pdf/"
	htmlOutDir   = "../../out/html/"
	readme       = "../../readme.md"
	testDataDir  = "../../testdata/markdown/"
	mdCheatsheet = testDataDir + "markdown-cheatsheet.md"
)

func runParse() {
	// render.BlackfridayPDF(pdfOutDir + "blackfriday.test.3.pdf")
	// render.GitHubPDF(pdfOutDir + "github8.pdf")

	// err := md2pdf.Md2PDF("/Users/clay/go/src/github.com/tcd/md2pdf/testdata/markdown/sample/blackfriday.readme.md", pdfOutDir+"blackfriday.pdf")
	err := md2pdf.Md2PDF("/Users/clay/Documents/VBI/Levys/levys-rails/meeting-with-joseph.md", pdfOutDir+"joseph.pdf")
	if err != nil {
		fmt.Println(err)
	}

	// err := md2pdf.Md2HTMLFile("/Users/clay/go/src/github.com/tcd/md2pdf/testdata/markdown/sample/blackfriday.readme.md", htmlOutDir+"blackfriday.html")
	// err := md2pdf.Md2PDF("/Users/clay/go/src/github.com/tcd/md2pdf/testdata/markdown/sample/blackfriday.readme.md", pdfOutDir+"blackfriday.2.pdf")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// parse.GenAST()

	// err := md2pdf.Md2HTMLFile(testDataDir+"list.md", htmlOutDir+"list.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = md2pdf.Md2PDF(testDataDir+"list.md", pdfOutDir+"list.pdf")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
