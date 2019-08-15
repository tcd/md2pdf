package md2pdf

var escapes = map[string]string{
	"&quot;":  `"`,
	"&amp;":   "&",
	"&lt;":    "<",
	"&gt;":    ">",
	"&rsquo;": "'",
}
