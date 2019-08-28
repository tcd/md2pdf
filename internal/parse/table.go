package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

// Table ...
func Table(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tableContent := parseTable(z)
	render.Table(pdf, tableContent)
}

func parseTable(z *html.Tokenizer) model.TableContent {
	tableContent := model.TableContent{}

	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "thead" {
				headers, alignments := parseTableHeaders(z)
				tableContent.AddRow(headers)
				tableContent.Alignments = alignments
			}
			if T1.Data == "tbody" {
				bodyRows := parseTableBody(z)
				tableContent.AddRows(bodyRows...)
			}
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "table" {
				break
			}
		}
	}

	return tableContent
}

func parseTableHeaders(z *html.Tokenizer) (headers, alignments []string) {
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "th" {
				content, alignment := parseTableHeaderRow(z, T1)
				headers = append(headers, content)
				alignments = append(alignments, alignment)
			}
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "tr" || T1.Data == "thead" {
				break
			}
		}
	}
	return headers, alignments
}

func parseTableHeaderRow(z *html.Tokenizer, startToken html.Token) (content, alignment string) {
	alignment = "L"

	for _, a := range startToken.Attr {
		if a.Key == "align" {
			if a.Val == "center" {
				alignment = "C"
			}
			if a.Val == "right" {
				alignment = "R"
			}
		}
	}

	for {
		tt := z.Next()
		if tt == html.TextToken {
			content = string(z.Text())
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "th" {
				break
			}
		}
	}
	return content, alignment
}

func parseTableBody(z *html.Tokenizer) [][]string {
	var rows [][]string
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "tr" {
				content := parseTableBodyRow(z)
				rows = append(rows, content)
			}
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "tbody" {
				break
			}
		}
	}
	return rows
}

func parseTableBodyRow(z *html.Tokenizer) []string {
	var columns []string
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "td" {
				content := parseTableBodyCell(z)
				columns = append(columns, content)
			}
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "tr" {
				break
			}
		}
	}
	return columns
}

func parseTableBodyCell(z *html.Tokenizer) string {
	var content string
	for {
		tt := z.Next()
		if tt == html.TextToken {
			column := string(z.Text())
			if column != "" {
				content = content + column
			}
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "td" {
				break
			}
		}
	}
	return content
}
