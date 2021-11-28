To execute test need to install docker and go

# UI

Steps to run from scratch UI test:
```
docker pull selenoid/video-recorder:latest-release
docker pull selenoid/vnc_chrome:93.0
docker-compose up -d
go mod tidy
cd ui
go test -v --godog.tags=RT
```

List of tags:
- RT - run full test, regression testing 
- from @one to @ten - run certain test 

To check selenoid console go to the:
http://localhost:8080

# API

Steps to run from scratch UI test:
```
go mod tidy
cd api
go test -v --godog.tags=API
```

List of tags:
- API - run full test
- from @one to @five - run certain test 