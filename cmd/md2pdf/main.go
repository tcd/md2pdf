package main

import (
	"github.com/tcd/md2pdf/md2pdf"
)

const (
	pdfOutDir  = "../../out/pdf/"
	htmlOutDir = "../../out/html/"
	readme     = "../../readme.md"
)

func main() {
	md2pdf.GitHubPDF(pdfOutDir + "github.pdf")
}
