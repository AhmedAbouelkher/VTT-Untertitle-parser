#!/bin/sh

mkdir -p ./bin

GOOS=windows go build -ldflags "-s -w" -o ./bin/vtt-parser.exe *.go &
    GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/linux_amd64_vtt-parser *.go &
        GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/darwin_amd64_vtt-parser *.go &
wait
echo "Build executables for windows, linux and macOS"

cp ./bin/darwin_amd64_vtt-parser /usr/local/bin/vtt-parser
echo "Copied vtt-parser to /usr/local/bin/vtt-parser"

chmod +x /usr/local/bin/vtt-parser
echo "Changed permissions to executable"

gzip -9 ./bin/vtt-parser.exe &
    gzip -9 ./bin/linux_amd64_vtt-parser &
        gzip -9 ./bin/darwin_amd64_vtt-parser &
wait
echo "Compressed executables into separate gz files"