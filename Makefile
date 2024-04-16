.PHONY: build run clean test

build:
	docker build -t spellchecker-app .

run: build
	docker run -p 8080:8080 --rm spellchecker-app

clean:
	docker rmi spellchecker-app

test:
	go test -v ./...
