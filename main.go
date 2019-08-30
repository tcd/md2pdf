package main

import (
	"log"

	"github.com/tcd/md2pdf/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
