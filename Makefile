PROJECT_NAME := "go-prntserve"
BINARY_NAME := "prntserve"
BINARIES_FOLDER := "builds"
GIT_SHA := $(shell git rev-parse HEAD)

dev: create_output_folder
	go build -ldflags '-X main.buildNumber=${GIT_SHA}' -o "${BINARIES_FOLDER}/dev/${BINARY_NAME}" .

release: create_output_folder
	go build -ldflags '-X main.buildNumber=${GIT_SHA} -X main.distribution=official' -o "${BINARIES_FOLDER}/releae/${BINARY_NAME}-release" .

clean:
	rm -rf ${BINARIES_FOLDER}
