## How To Run
### Run HTTP & GRPC Server
``
cd q2
go run cmd/main.go
``

### Run GRPC Client
```
cd q2
go run cmd/grpc_client.go
```

### APIs
#### Search Movies
`/movies/search?pagination=1&searchword=batman`

#### Get Single Movie
`/movie?id={id}`