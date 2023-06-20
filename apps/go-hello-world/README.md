# Kindle GUI App Hello World

## Libraries
`nucular` is the GUI library. Commit hash `1370bfe9ca4c` is used as it's from around when `github.com/efskap/shiny-kindle` was created. Newer versions have a dependency issue. FixTBD.

```
require github.com/aarzilli/nucular v0.0.0-20200502123933-1370bfe9ca4c

replace golang.org/x/exp => github.com/efskap/shiny-kindle v0.0.0-20200523223607-07a3b6aac3a9
```

## Setup
Generate the `go.sum` file thing:
```bash
go mod tidy
```

## Building
`GOARCH` is set to *arm*, because the kindle is an ARM device
`GOARM` is the archatecture. See https://zchee.github.io/golang-wiki/GoArm/
```bash
GOARCH=arm GOARM=5 go build
```

## Running
```bash
# on the kindle
## stop the Kindle GUI.
stop lab126_gui
## start the binary
DISPLAY=:0 /mnt/us/kindle_test
```
