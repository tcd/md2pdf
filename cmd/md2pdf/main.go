package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func init() {
	flag.Usage = usage
}

func main() {
	flag.Parse()
}

func usage() {
	fmt.Println(titleString())
	fmt.Printf("Usage: %s [OPTIONS] argument\n", os.Args[0])
	flag.PrintDefaults()
}

func titleString() string {
	lines := []string{
		"		_  _____           _  __",
		"		| |/ __  \\         | |/ _|",
		"   _ __ ___   __| |`' / /'_ __   __| | |_",
		"  | '_ ` _ \\ / _` |  / / | '_ \\ / _` |  _|",
		"  | | | | | | (_| |./ /__| |_) | (_| | |",
		"  |_| |_| |_|\\__,_|\\_____/ .__/ \\__,_|_|",
		"			 | |",
		"			 |_|",
		"",
	}
	return strings.Join(lines, "\n")
}
