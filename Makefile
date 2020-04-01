.PHONY all deps run clean docker docker-build docker-push docker-test

PYTHON = python
DOCKER = docker

all: deps run

deps:
  pip install -f requirements.txt

run:
  python ./main.py

clean:
  rm -rf *.pyc

docker: docker-build docker-push

docker-test:
   test $(DOCKERREPO)

docker-build: docker-test
   docker build . $(DOCKERREPO)

docker-push: docker-test
   docker push $(DOCKERREPO)