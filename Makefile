.PHONY: build install build-unversioned

build:
	go build -ldflags "-X github.com/zuma206/sysmig/utils.VERSION=`git describe --tags --abbrev=0`"

install:
	sudo cp ./sysmig /usr/local/bin

build-unversioned:
	go build