#!/bin/sh

mkdir -p ./bin

GOOS=windows go build -ldflags "-s -w" -o ./bin/vvt-parser.exe *.go &
    GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/linux_amd64_vvt-parser *.go &
        GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/darwin_amd64_vvt-parser *.go &
wait
echo "Build executables for windows, linux and macOS"

cp ./bin/darwin_amd64_vvt-parser /usr/local/bin/vvt-parser
echo "Copied vvt-parser to /usr/local/bin/vvt-parser"

chmod +x /usr/local/bin/vvt-parser
echo "Changed permissions to executable"

gzip -9 ./bin/vvt-parser.exe &
    gzip -9 ./bin/linux_amd64_vvt-parser &
        gzip -9 ./bin/darwin_amd64_vvt-parser &
wait
echo "Compressed executables into separate gz files"