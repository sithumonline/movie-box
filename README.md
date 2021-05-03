# Torrent-Box

Download movie form [YTS](https://yts.mx/) without visiting to [YTS](https://yts.mx/api#list_movies)

Built top on [anacrolix/torrent](https://github.com/anacrolix/torrent) lib

## Help

```
movie-box -h
```

## Start download

```
movie-box get -n "name of the movie" -q 1080p -o "/path/to/download/directory"
```

## With docker

```
docker run --rm -it ghcr.io/sithumonline/movie-box:0.0.2 -n "name of the movie" -q 720p -o "/path/to/download/directory"
```
