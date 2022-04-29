PROJECT_NAME := "go-prntserve"
BINARY_NAME := "prntserve"
BINARIES_FOLDER := "builds"
GIT_SHA := $(shell git rev-parse HEAD)

dev:
	go build -ldflags '-X main.buildNumber=${GIT_SHA}' -o "${BINARIES_FOLDER}/dev/${BINARY_NAME}" .

release:
	# Thanks go to https://github.com/lawl/NoiseTorch/blob/master/Makefile
	go build -trimpath -tags release -a -ldflags '-s -w -extldflags "-static" -X main.buildNumber=${GIT_SHA} -X main.distribution=official' -o "${BINARIES_FOLDER}/release/${BINARY_NAME}" .

clean:
	rm -rf ${BINARIES_FOLDER}
