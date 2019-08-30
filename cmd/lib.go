package cmd

import (
	"log"
	"strings"
)

// log.Fatal if the error isn't nil.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// return a nice Ascii title for "md2pdf"
var titleString = strings.Join(titleLines, "\n")
var titleLines = []string{
	"               _  _____           _  __",
	"              | |/ __  \\         | |/ _|",
	" _ __ ___   __| |`' / /'_ __   __| | |_",
	"| '_ ` _ \\ / _` |  / / | '_ \\ / _` |  _|",
	"| | | | | | (_| |./ /__| |_) | (_| | |",
	"|_| |_| |_|\\__,_|\\_____/ .__/ \\__,_|_|",
	"                       | |",
	"                       |_|",
}

var titleLogo = `(\\
\\'\\
 \\'\\     __________
 / '|   ()_________)
 \\ '/    \\ ~~~~~~~~ \\
   \\       \\ ~~~~~~   \\
   ==).      \\__________\\
  (__)       ()__________)`
