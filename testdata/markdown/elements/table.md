
## Tables

Tables aren't part of the core Markdown spec, but they are part of GFM and *Markdown Here* supports them. They are an easy way of adding tables to your email -- a task that would otherwise require copy-pasting from another application.

```no-highlight
Colons can be used to align columns.

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

There must be at least 3 dashes separating each header cell.
The outer pipes (|) are optional, and you don't need to make the 
raw Markdown line up prettily. You can also use inline Markdown.

Markdown | Less | Pretty
--- | --- | ---
*Still* | `renders` | **nicely**
1 | 2 | 3
```

Colons can be used to align columns.

| Tables        | Are           | Cool |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

There must be at least 3 dashes separating each header cell. The outer pipes (|) are optional, and you don't need to make the raw Markdown line up prettily. You can also use inline Markdown.

Markdown | Less | Pretty
--- | --- | ---
*Still* | `renders` | **nicely**
1 | 2 | 3

Here's something to squish between tables.

| command                       | summary                                                     |
|-------------------------------|-------------------------------------------------------------|
| `ls`                          | list names of all files in current directory                |
| `ls` *filenames*              | list only the named files                                   |
| `ls -t`                       | list in time order, most recent first                       |
| `ls -l`                       | list long: more information; also `ls -lt`                  |
| `ls -u`                       | list by last time used; also `ls -lu`, `ls -lut`            |
| `ls -r`                       | list in reverse order; also `-rt`, `rlt`, etc.              |
| `ed` *filename*               | edit named file                                             |
| `cp` *file1 file2*            | copy *file1* to *file2*, overwrite old *file2* if it exists |
| `mv` *file1 file2*            | move *file1* to *file2*, overwrite old *file2* if it exists |
| `rm` *filenames*              | remove named files, irrevocably                             |
| `cat` *filenames*             | print contents of named files                               |
| `wc` *filenames*              | count lines, words, and characters for each file            |
| `wc -l` *filenames*           | count lines for each file                                   |
| `grep` *pattern filenames*    | print lines matching *pattern*                              |
| `grep -v` *pattern filenames* | print lines not matching *pattern*                          |
| `sort` *filenames*            | sort files alphabetically by line                           |
| `tail` *filename*             | print last 10 lines of file                                 |
| `tail -n` *filename*          | print last *n* lines of file                                |
| `tail +n` *filename*          | start printing file at line *n*                             |
| `cmp` *file1 file2*           | print location of first difference                          |
| `diff` *file1 file2*          | print all differences between files                         |


And here's a nice paragraph to round things out.


## Chroma's Supported languages

Prefix | Language
:----: | --------
A | ABNF, ActionScript, ActionScript 3, Ada, Angular2, ANTLR, ApacheConf, APL, AppleScript, Arduino, Awk
B | Ballerina, Base Makefile, Bash, Batchfile, BlitzBasic, BNF, Brainfuck
C | C, C#, C++, Cap'n Proto, Cassandra CQL, Ceylon, CFEngine3, cfstatement, ChaiScript, Cheetah, Clojure, CMake, COBOL, CoffeeScript, Common Lisp, Coq, Crystal, CSS, Cython
D | D, Dart, Diff, Django/Jinja, Docker, DTD
E | EBNF, Elixir, Elm, EmacsLisp, Erlang
F | Factor, Fish, Forth, Fortran, FSharp
G | GAS, GDScript, Genshi, Genshi HTML, Genshi Text, GLSL, Gnuplot, Go, Go HTML Template, Go Text Template, GraphQL, Groovy
H | Handlebars, Haskell, Haxe, HCL, Hexdump, HTML, HTTP, Hy
I | Idris, INI, Io
J | Java, JavaScript, JSON, Julia, Jungle
K | Kotlin
L | Lighttpd configuration file, LLVM, Lua
M | Mako, markdown, Mason, Mathematica, Matlab, MiniZinc, Modula-2, MonkeyC, MorrowindScript, Myghty, MySQL
N | NASM, Newspeak, Nginx configuration file, Nim, Nix
O | Objective-C, OCaml, Octave, OpenSCAD, Org Mode
P | PacmanConf, Perl, PHP, Pig, PkgConfig, PL/pgSQL, plaintext, PostgreSQL SQL dialect, PostScript, POVRay, PowerShell, Prolog, Protocol Buffer, Puppet, Python, Python 3
Q | QBasic
R | R, Racket, Ragel, react, reg, reStructuredText, Rexx, Ruby, Rust
S | Sass, Scala, Scheme, Scilab, SCSS, Smalltalk, Smarty, Snobol, Solidity, SPARQL, SQL, SquidConf, Swift, SYSTEMD, systemverilog
T | TASM, Tcl, Tcsh, Termcap, Terminfo, Terraform, TeX, Thrift, TOML, TradingView, Transact-SQL, Turing, Turtle, Twig, TypeScript, TypoScript, TypoScriptCssData, TypoScriptHtmlData
V | VB.net, verilog, VHDL, VimL, vue
W | WDTE
X | XML, Xorg
Y | YAML

## Now Reversed

hmm... what if we have some text here?

|                                                                                     Language                                                                                     | Prefix |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----: |
| ABNF, ActionScript, ActionScript 3, Ada, Angular2, ANTLR, ApacheConf, APL, AppleScript, Arduino, Awk                                                                             |   A    |
| Ballerina, Base Makefile, Bash, Batchfile, BlitzBasic, BNF, Brainfuck                                                                                                            |   B    |
| C, C#, C++, Cap'n Proto, Cassandra CQL, Ceylon, CFEngine3, cfstatement, ChaiScript, Cheetah, Clojure, CMake, COBOL, CoffeeScript, Common Lisp, Coq, Crystal, CSS, Cython         |   C    |
| D, Dart, Diff, Django/Jinja, Docker, DTD                                                                                                                                         |   D    |
| EBNF, Elixir, Elm, EmacsLisp, Erlang                                                                                                                                             |   E    |
| Factor, Fish, Forth, Fortran, FSharp                                                                                                                                             |   F    |
| GAS, GDScript, Genshi, Genshi HTML, Genshi Text, GLSL, Gnuplot, Go, Go HTML Template, Go Text Template, GraphQL, Groovy                                                          |   G    |
| Handlebars, Haskell, Haxe, HCL, Hexdump, HTML, HTTP, Hy                                                                                                                          |   H    |
| Idris, INI, Io                                                                                                                                                                   |   I    |
| Java, JavaScript, JSON, Julia, Jungle                                                                                                                                            |   J    |
| Kotlin                                                                                                                                                                           |   K    |
| Lighttpd configuration file, LLVM, Lua                                                                                                                                           |   L    |
| Mako, markdown, Mason, Mathematica, Matlab, MiniZinc, Modula-2, MonkeyC, MorrowindScript, Myghty, MySQL                                                                          |   M    |
| NASM, Newspeak, Nginx configuration file, Nim, Nix                                                                                                                               |   N    |
| Objective-C, OCaml, Octave, OpenSCAD, Org Mode                                                                                                                                   |   O    |
| PacmanConf, Perl, PHP, Pig, PkgConfig, PL/pgSQL, plaintext, PostgreSQL SQL dialect, PostScript, POVRay, PowerShell, Prolog, Protocol Buffer, Puppet, Python, Python 3            |   P    |
| QBasic                                                                                                                                                                           |   Q    |
| R, Racket, Ragel, react, reg, reStructuredText, Rexx, Ruby, Rust                                                                                                                 |   R    |
| Sass, Scala, Scheme, Scilab, SCSS, Smalltalk, Smarty, Snobol, Solidity, SPARQL, SQL, SquidConf, Swift, SYSTEMD, systemverilog                                                    |   S    |
| TASM, Tcl, Tcsh, Termcap, Terminfo, Terraform, TeX, Thrift, TOML, TradingView, Transact-SQL, Turing, Turtle, Twig, TypeScript, TypoScript, TypoScriptCssData, TypoScriptHtmlData |   T    |
| VB.net, verilog, VHDL, VimL, vue                                                                                                                                                 |   V    |
| WDTE                                                                                                                                                                             |   W    |
| XML, Xorg                                                                                                                                                                        |   X    |
| YAML                                                                                                                                                                             |   Y    |

## `gofpdf` Files

|        file         |                                        description                                        |                notes                 |
| ------------------- | ----------------------------------------------------------------------------------------- | ------------------------------------ |
| `list/list.go`      |                                                                                           | Unused?                              |
| `compare.go`        | Functions for comparing Bytes & Pdfs                                                      | Unused?                              |
| `def.go`            | Variable and Type declarations                                                            |                                      |
| `doc.go`            | package documentation for `godoc`                                                         |                                      |
| `embedded.go`       | Data for embedded standard fonts                                                          |                                      |
| `font.go`           | Functions for parsing and embedding font files in PDFs                                    | `MakeFont()` is defined in this file |
| `fpdf.go`           | Main file                                                                                 |                                      |
| `fpdf_test.go`      | *tests*                                                                                   | test                                 |
| `fpdftrans.go`      | Functions for transforming, translating, & scaling PDF content (text, drawings, & images) |                                      |
| `grid.go`           | Utility methods for drawing graphs                                                        |                                      |
| `htmlbasic.go`      | Interface for drawing basic HTML to a PDF                                                 |                                      |
| `label.go`          | Functions used in `grid.go` for labeling graphs                                           |                                      |
| `layer.go`          | Methods adding [PDF layer][1] functionality                                               |                                      |
| `png.go`            | PNG parsing functions                                                                     |                                      |
| `protect.go`        | Functions for encrypting & password protecting PDF files                                  |                                      |
| `splittext.go`      | `fpdf.SplitText()` Definition                                                             | Single Function                      |
| `spotcolor.go`      | Methods enabling the use of [spot colors][2] (used in professional printing)              |                                      |
| `subwrite.go`       | Method for writing subscripts and superscripts                                            | Single Function                      |
| `svgbasic.go`       | Super basic SVG parsing                                                                   |                                      |
| `svgwrite.go`       | Super basic SVG rendering                                                                 | Single Function                      |
| `template.go`       |                                                                                           |                                      |
| `template_impl.go`  |                                                                                           |                                      |
| `ttfparser.go`      | Utility to parse TTF font files                                                           |                                      |
| `ttfparser_test.go` | *tests*                                                                                   | test                                 |
| `utf8fontfile.go`   |                                                                                           |                                      |
| `util.go`           |                                                                                           |                                      |
