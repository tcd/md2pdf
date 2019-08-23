package render

import (
	"regexp"
	"strings"
)

// https://www.w3.org/wiki/Common_HTML_entities_used_for_typography
var escapes = map[string]string{
	"&quot;": `"`,
	"&amp;":  "&",
	"&lt;":   "<",
	"&gt;":   ">",

	"&#39;": "'",
	"&#34;": `"`,
	// "&rdquo;":  `"`,    // ”
	// "&ldquo;":  `"`,    // “
	// "&rsquo;":  "'",    // ’
	// "&lsquo;":  "'",    // ‘
	// "&mdash;":  "--",   // —
	"—": "-",
	// "&ndash;":  "-",    // –
	// "&copy;":   "(c)",  // ©
	// "&trade;":  "(tm)", // ™
	// "&reg;":    "(r)",  // ®
	// "&hellip;": "...",  // …
	// "&frac12;": "1/2",  // ½
	// "&frac14;": "1/4",  // ¼
	// "&frac34;": "3/4",  // ¾
}

func cleanString(str string) string {
	newStr := str[:]
	for k, v := range escapes {
		newStr = strings.ReplaceAll(newStr, k, v)
	}
	re := regexp.MustCompile(`\t+`)
	newStr = re.ReplaceAllString(newStr, " ")
	return newStr
}
