# md2pdf

<!-- [![Golang Documentation](https://godoc.org/github.com/tcd/md2pdf?status.svg)](https://godoc.org/github.com/tcd/md2pdf) -->
[![GitHub issues](https://img.shields.io/github/issues/tcd/md2pdf.svg)](https://github.com/tcd/md2pdf/issues)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftcd%2Fmd2pdf.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftcd%2Fmd2pdf?ref=badge_shield)
[![License](https://img.shields.io/github/license/tcd/md2pdf.svg)](https://github.com/tcd/md2pdf/blob/master/LICENSE)


![fancy-gif](https://raw.githubusercontent.com/dunstontc/assets/master/gifs/fanciness.gif)

## About

Parse markdown files and output PDFs that resemble GitHub's markdown rendering.

### What's Supported?

- Paragraphs
- Blockquotes
- Codeblocks
    - Syntax highlighting provided by [Chroma](https://github.com/alecthomas/chroma)
- All kinds of emphasis:
    - **Bold**
    - *Italic*
    - `Code`
    - ~~Strike~~
    - [Links](https://github.com/tcd/md2pdf)
- Images
    - jpeg, png, and gif\*
      - Only the first frame of a gif is drawn to the pdf :pensive:
- Headers
- Horizontal Rules
- Links
    - Links on images
    - External Links
- Tables

### What Isn't Supported?

- No Emoji 😔
- No SVG 
- No Inline HTML
- No Definition Lists
- No Task Lists

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
