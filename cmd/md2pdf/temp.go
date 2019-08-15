package main

import (
	"fmt"

	"github.com/tcd/md2pdf/md2pdf"
)

const (
	pdfOutDir    = "../../out/pdf/"
	htmlOutDir   = "../../out/html/"
	readme       = "../../readme.md"
	mdCheatsheet = "../../static/markdown-cheatsheet.md"
)

func runParse() {
	err := md2pdf.Md2PDF(mdCheatsheet, pdfOutDir+"recursive3.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

func cleanHTML() {
	err := md2pdf.Md2HTMLFile(mdCheatsheet, htmlOutDir+"cleaner.html")
	if err != nil {
		fmt.Println(err)
	}
}
