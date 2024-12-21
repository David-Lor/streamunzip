# StreamUnzip

Utility for downloading and extracting the contents of a ZIP file directly.

When downloading a ZIP file from anywhere, we normally download the ZIP, then extract the file/s we want, then delete the ZIP.
This utility downloads the extracted contents from that ZIP file directly, skipping this step.
This can be useful when dealing with very big files, avoiding disk usage.

## Installing/Running

Run as a normal Golang app (`cd src && go run .`). Developed and tested with Go version 1.21.5 under Linux, but should work under any OS.

## Usage

At this moment this utility only works via terminal with an interactive mode. It asks the ZIP URL and extraction path read from the keyboard input (stdin).

Then, the tool will ask what to do for each file:

- Save the shown file
- Skip this file
- Save this and all the next files, not asking for any more confirmations
- Quit

Since the read of a ZIP file is sequential, if you skip a file, that file will still be downloaded to your system, but its contents ignored until the next file.

### Example

```
$ cd src
$ go run .

Enter ZIP URL to download: https://github.com/David-Lor/pvpc/archive/refs/heads/main.zip
Enter download path: /tmp/pvpc-extract
What to do with this file?
        Name: pvpc-main/.github/workflows/export-data-range.yaml
        Size: 3.1 kB
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:44 Skip file pvpc-main/.github/workflows/export-data-range.yaml
What to do with this file?
        Name: pvpc-main/.github/workflows/export-data.yaml
        Size: 4.7 kB
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:47 Skip file pvpc-main/.github/workflows/export-data.yaml
What to do with this file?
        Name: pvpc-main/.github/workflows/update-license-year.yaml
        Size: 486 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:48 Skip file pvpc-main/.github/workflows/update-license-year.yaml
What to do with this file?
        Name: pvpc-main/.gitignore
        Size: 51 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:49 Skip file pvpc-main/.gitignore
What to do with this file?
        Name: pvpc-main/LICENSE.md
        Size: 755 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:50 Skip file pvpc-main/LICENSE.md
What to do with this file?
        Name: pvpc-main/README.md
        Size: 1.9 kB
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:51 Skip file pvpc-main/README.md
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/01.json
        Size: 481 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:53 Saving file pvpc-main/data/cm/2021/06/01.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/02.json
        Size: 484 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:54 Saving file pvpc-main/data/cm/2021/06/02.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/03.json
        Size: 486 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:55 Saving file pvpc-main/data/cm/2021/06/03.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/04.json
        Size: 486 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:56 Saving file pvpc-main/data/cm/2021/06/04.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/05.json
        Size: 484 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:57 Skip file pvpc-main/data/cm/2021/06/05.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/06.json
        Size: 483 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:57 Skip file pvpc-main/data/cm/2021/06/06.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/07.json
        Size: 483 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:58 Skip file pvpc-main/data/cm/2021/06/07.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/08.json
        Size: 485 B
[s]Save [k]Skip [a]Save all [q]Quit
k
2024/12/21 16:55:58 Skip file pvpc-main/data/cm/2021/06/08.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/09.json
        Size: 483 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:59 Saving file pvpc-main/data/cm/2021/06/09.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/10.json
        Size: 485 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:59 Saving file pvpc-main/data/cm/2021/06/10.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/11.json
        Size: 487 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:59 Saving file pvpc-main/data/cm/2021/06/11.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/12.json
        Size: 483 B
[s]Save [k]Skip [a]Save all [q]Quit
s
2024/12/21 16:55:59 Saving file pvpc-main/data/cm/2021/06/12.json
What to do with this file?
        Name: pvpc-main/data/cm/2021/06/13.json
        Size: 486 B
[s]Save [k]Skip [a]Save all [q]Quit
q
2024/12/21 16:56:00 Quit


$ ls /tmp/pvpc-extract/pvpc-main/data/cm/2021/06/
01.json  02.json  03.json  04.json  09.json  10.json  11.json  12.json
```

## Known issues

- The server may close the connection if you take too long to choose an answer for a file.

## Changelog

- v0.1
  - Initial release: main functionality, interactive terminal mode.
