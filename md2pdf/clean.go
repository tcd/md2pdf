package md2pdf

import (
	"strings"
)

// https://www.w3.org/wiki/Common_HTML_entities_used_for_typography
var escapes = map[string]string{
	"&quot;": `"`,
	"&amp;":  "&",
	"&lt;":   "<",
	"&gt;":   ">",
	// "&rdquo;":  `"`,
	// "&rsquo;":  "'",
	// "&mdash;":  "",
	// "&ndash;":  "",
	// "&copy;":   "",
	// "&trade;":  "",
	// "&reg;":    "",
	// "&hellip;": "",
	// "&frac12;": "",
	// "&frac14;": "",
	// "&frac34;": "",
}

func cleanHTML(html string) string {
	s1 := strings.ReplaceAll(html, "&quot;", `"`)
	s2 := strings.ReplaceAll(s1, "&amp;", "&")
	s3 := strings.ReplaceAll(s2, "&lt;", "<")
	s4 := strings.ReplaceAll(s3, "&gt;", ">")
	return s4

	// r := strings.NewReplacer( // TODO: File an issue about strings.NewReplacer. It missed some '&lt;'s and '&gt;'s.
	// 	"&quot;", `"`,
	// 	"&amp;", "&",
	// 	"&lt;", "<",
	// 	"&gt;", ">",
	// )
	// return r.Replace(html)
}
