FILENAME=webserve

all: clean deps build
deps:
	go get github.com/valyala/fasthttp
	go get github.com/joho/godotenv

build:
	go build -o $(FILENAME) -v

clean:
	go clean
	rm -f $(filename)