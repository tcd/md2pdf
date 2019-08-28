package parse

import (
	"github.com/tcd/md2pdf/internal/renderable"
	"golang.org/x/net/html"
)

// Image gathers the data needed to render an image.
func Image(token html.Token) renderable.Image {
	return renderable.Image{
		Src: parseImg(token),
	}
}

func parseImg(token html.Token) string {
	var src string
	for _, a := range token.Attr {
		if a.Key == "src" {
			src = a.Val
		}
	}
	return src
}
