package parse

import (
	renderer "github.com/tcd/md2pdf/internal/r1"
	"golang.org/x/net/html"
)

// Image gathers the data needed to render an image.
func Image(token html.Token) renderer.Image {
	return renderer.Image{
		Type: "image",
		Src:  parseImg(token),
	}
}

// LinkedImage gathers the data needed to render an image with an embedded link.
func LinkedImage(token html.Token, href string) renderer.Image {
	return renderer.Image{
		Type: "image",
		Src:  parseImg(token),
		Link: href,
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
