# md2pdf

## Markdown Specs

- [GitHub Flavored Markdown](https://github.github.com/gfm/)

## Dependencies

### Parsing

- [russross/blackfriday](https://github.com/russross/blackfriday) - Blackfriday: a markdown processor for Go
  - [Depado/bfchroma](https://github.com/Depado/bfchroma) - Integrating Chroma syntax highlighter as a Blackfriday renderer 
  - [shurcooL/github_flavored_markdown](https://github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links.

### Rendering

- [jung-kurt/gofpdf](https://github.com/jung-kurt/gofpdf) - A PDF document generator with high level support for text, drawing and images
- Usage Examples
  - [ajstarks/deck](https://github.com/ajstarks/deck/blob/master/cmd/pdfdeck/pdfdeck.go)
  - [sprucehealth/gofpdf](https://godoc.org/github.com/sprucehealth/gofpdf)
  - [johnfercher/maroto](https://github.com/johnfercher/maroto)

### Styling

- [sindresorhus/github-markdown-css](https://github.com/sindresorhus/github-markdown-css) - The minimal amount of CSS to replicate the GitHub Markdown style
