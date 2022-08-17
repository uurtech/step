BINARY_NAME=step

build:
	go install
	go build

run:
	./${BINARY_NAME}

build_and_run: build run

install:
	make build
	sudo mv ./${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

clean:
	go clean
	