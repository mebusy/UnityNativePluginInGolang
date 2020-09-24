# to list all GOOS/GOARCH:  go tool dist list

# c-shared

# OSX
DIST="../dist/OSX/"
mkdir -p $DIST
echo $DIST
GOOS=darwin GOARCH=amd64 go build -v -ldflags "-w" -buildmode=c-archive -o $DIST/libgo.a  libgo/

# Win64
DIST="../dist/Win64/"
echo $DIST
mkdir -p $DIST
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++   go build -v -ldflags "-w" -buildmode=c-shared -o $DIST/libgo.dll  libgo/
