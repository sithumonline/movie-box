rm -rf build/*
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/movie-box_0.0.1_linux_amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-w -s" -o build/movie-box_0.0.1_linux_arm .
docker buildx build --platform linux/amd64,linux/arm/v7 -t ghcr.io/sithumonline/movie-box:0.0.4 \
      --cache-to=type=registry,ref=ghcr.io/sithumonline/movie-box:0.0.4-cache \
      --cache-from=type=registry,ref=ghcr.io/sithumonline/movie-box:0.0.4-cache \
      --push .
