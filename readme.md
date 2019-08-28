# md2pdf

<!-- [![Golang Documentation](https://godoc.org/github.com/tcd/md2pdf?status.svg)](https://godoc.org/github.com/tcd/md2pdf) -->
[![GitHub issues](https://img.shields.io/github/issues/tcd/md2pdf.svg)](https://github.com/tcd/md2pdf/issues)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftcd%2Fmd2pdf.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftcd%2Fmd2pdf?ref=badge_shield)
[![License](https://img.shields.io/github/license/tcd/md2pdf.svg)](https://github.com/tcd/md2pdf/blob/master/LICENSE)

Generate PDFs from Markdown files on the command line.
Small, fast, and free.

![fancy-gif](https://raw.githubusercontent.com/dunstontc/assets/master/gifs/fanciness.gif)

## Installation

### Building from Source

```bash
go get -u github.com/tcd/md2pdf
cd $GOPATH/github.com/tcd/md2pdf
make install # GO111MODULE=on go install
```


## Usage

```sh
# will generate markdown-file.pdf in your current directory.
md2pdf path/to/markdown-file.md
```

## Related Projects

- [mandolyte/mdtopdf](https://github.com/mandolyte/mdtopdf) - Markdown to PDF 
- [ajstarks/deck](https://github.com/ajstarks/deck) - Slide Decks
- [johnfercher/maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs.

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftcd%2Fmd2pdf.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftcd%2Fmd2pdf?ref=badge_large)
