pdfOutDir="out/pdf"
htmlOutDir="out/html"
readme="readme.md"
markdownDataDir="testdata/markdown"
htmlDataDir="testdata/html"
mdCheatsheet="$markdownDataDir/markdown-cheatsheet.v2.md"

# go run main.go "$readme"
go run main.go "$readme" --html
go run main.go "$readme"
