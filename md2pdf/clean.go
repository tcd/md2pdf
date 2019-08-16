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
	// "&rdquo;":  `"`,    // ”
	// "&ldquo;":  `"`,    // “
	// "&rsquo;":  "'",    // ’
	// "&lsquo;":  "'",    // ‘
	// "&mdash;":  "--",   // —
	// "&ndash;":  "-",    // –
	// "&copy;":   "(c)",  // ©
	// "&trade;":  "(tm)", // ™
	// "&reg;":    "(r)",  // ®
	// "&hellip;": "...",  // …
	// "&frac12;": "1/2",  // ½
	// "&frac14;": "1/4",  // ¼
	// "&frac34;": "3/4",  // ¾
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
