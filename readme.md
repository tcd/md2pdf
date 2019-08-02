# md2pdf

## Markdown Specs

- [GitHub Flavored Markdown](https://github.github.com/gfm/)

## Dependencies

- Parsing
- [russross/blackfriday](https://github.com/russross/blackfriday) - Blackfriday: a markdown processor for Go
- [Depado/bfchroma](https://github.com/Depado/bfchroma) - Integrating Chroma syntax highlighter as a Blackfriday renderer 
- Rendering
- [jung-kurt/gofpdf](https://github.com/jung-kurt/gofpdf) - A PDF document generator with high level support for text, drawing and images
  - Examples
    - [ajstarks/deck](https://github.com/ajstarks/deck/blob/master/cmd/pdfdeck/pdfdeck.go)
- Styling
  - [sindresorhus/github-markdown-css](https://github.com/sindresorhus/github-markdown-css) - The minimal amount of CSS to replicate the GitHub Markdown style
