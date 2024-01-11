# Build windows .exe
default:
    #!/usr/bin/env bash
    source ~/.profilego
    go get -u
    GOOS=windows GOARCH=amd64 go build