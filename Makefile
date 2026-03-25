APP_NAME := writersite

.PHONY: build run clean

build:
	go build -o cmd/$(APP_NAME)

run: build
	./cmd/$(APP_NAME)

clean:
	rm -rf bin

run-ui:
	npm --prefix ./ui run dev
