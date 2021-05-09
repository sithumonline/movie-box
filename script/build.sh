rm -rf build/*
export VERSION=0.0.4
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/movie-box_"$VERSION"_linux_amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-w -s" -o build/movie-box_"$VERSION"_linux_arm .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o build/movie-box_"$VERSION"_windows_amd64 .
docker buildx build --platform linux/amd64,linux/arm/v7 -t ghcr.io/sithumonline/movie-box:$VERSION \
      --cache-to=type=registry,ref=ghcr.io/sithumonline/movie-box:$VERSION-cache \
      --cache-from=type=registry,ref=ghcr.io/sithumonline/movie-box:$VERSION-cache \
      --push .
