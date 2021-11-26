To execute test need to install docker and go

steps to run from scratch:
```
docker pull selenoid/video-recorder:latest-release
docker pull selenoid/vnc_chrome:93.0
docker-compose up -d
go mod tidy
cd src
go test -v --godog.tags=RT
```

List of tags:
- RT - run full test, regression testing 
- from @one to @ten - run certain test 

To check selenoid console go to the:
http://localhost:8080