CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-w -s" -o build/movie-box_0.0.1_linux_386 .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/movie-box_0.0.1_linux_amd64 .
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-w -s" -o build/movie-box_0.0.1_windows_386 .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o build/movie-box_0.0.1_windows_amd64 .
