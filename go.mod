module github.com/tcd/md2pdf

go 1.12

require (
	github.com/alecthomas/chroma v0.6.6
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/tcd/gofpdf-1 v0.1.15-0.20190910004141-2b18ef17504b
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
