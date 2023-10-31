BINARY_NAME = quiz_game
.DEFAULT_GOAL := run

run: build
	@./bin/$(BINARY_NAME) $(CSV)

build:
	@GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_NAME)

clean:
	@rm -f ./bin/$(BINARY_NAME)