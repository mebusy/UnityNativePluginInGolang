# to list all GOOS/GOARCH:  go tool dist list

# c-shared


# OSX
DIST="../dist/OSX/"
mkdir -p $DIST
GOOS=darwin GOARCH=amd64 go build -buildmode=c-archive -o $DIST/  libgo.go

