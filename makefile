BIN_PATH=./
BINARY=test
MAIN_FILE=main.go
VERSION=`git rev-parse --short HEAD`
BUILD=`date +%Y-%m-%d^%H:%M:%S`
LDFLAGS=-ldflags "-X xframe/config.BuildVersion=${VERSION} -X xframe/config.BuildAt=${BUILD}"

build:
	@go build ${LDFLAGS} -o ${BIN_PATH}/${BINARY} ${MAIN_FILE}

