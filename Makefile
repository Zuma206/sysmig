.PHONY: all

all:
	go build -ldflags "-X github.com/zuma206/sysmig/utils.VERSION=`git describe --tags --abbrev=0`"