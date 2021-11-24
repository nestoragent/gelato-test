
```
docker pull selenoid/vnc_chrome:93.0
docker-compose up -d
go mod tidy
cd src
go test -v --godog.tags=one
```