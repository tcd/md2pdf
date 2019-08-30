
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
