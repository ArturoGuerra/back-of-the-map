.PHONY: all clean build docker docker-build docker-push docker-test

GOBUILD = go build
APPNAME = rolewatcher

all: build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/main.go

docker-test:
	test $(DOCKERREPO)

docker-build: docker-test
	docker build . -t $(DOCKERREPO)

docker-push: docker-test
	docker push $(DOCKERREPO)

docker: docker-build docker-push