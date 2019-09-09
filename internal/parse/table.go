package parse

import (
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/renderer"
	"golang.org/x/net/html"
)

// Table gathers the data needed to render a table.
func Table(z *html.Tokenizer) renderer.Table {
	return renderer.Table{
		Type:    "table",
		Content: parseTable(z),
	}
}

func parseTable(z *html.Tokenizer) model.TableContent {
	tableContent := model.TableContent{}

	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "thead" {
				headers, alignments := parseTableHeaders(z)
				tableContent.AddRows(headers)
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

func parseTableHeaders(z *html.Tokenizer) (headers []model.Contents, alignments []string) {
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "th" {
				content, alignment := parseTableHeaderRow(z, T1)
				cell := model.Contents{}
				cell.AddStr(content)
				headers = append(headers, cell)
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

func parseTableBody(z *html.Tokenizer) [][]model.Contents {
	var rows [][]model.Contents
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "tr" {
				row := parseTableBodyRow(z)
				rows = append(rows, row)
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

func parseTableBodyRow(z *html.Tokenizer) []model.Contents {
	var columns []model.Contents
	for {
		tt := z.Next()
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "td" {
				cell := parseTableBodyCell(z)
				columns = append(columns, cell)
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

func parseTableBodyCell(z *html.Tokenizer) model.Contents {
	var cell model.Contents
	for {
		tt := z.Next()
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "td" {
				break
			}
		}
		if tt == html.TextToken {
			txt := string(z.Text())
			if txt != "" {
				cell.AddStr(txt)
			}
		}
		if tt == html.StartTagToken {
			T1 := z.Token()
			blankContent := model.Text{}
			parseContent(z, T1, blankContent, &cell)
		}
	}
	return cell
}
