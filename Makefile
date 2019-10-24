NAME = vault-cred-helper
VERSION = $(shell grep 'const version' main.go | cut -d'"' -f 2)

BUILD_CMD = CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo

clean:
	rm -rf $(NAME) ./dist
	go clean -i ./...

dist: clean
	mkdir ./dist

	GOOS=darwin $(BUILD_CMD) -o ./dist/$(NAME) .
	cd ./dist && tar -czf $(NAME)-darwin-$(VERSION).tgz $(NAME); rm -f $(NAME)

	GOOS=linux $(BUILD_CMD) -o ./dist/$(NAME) .
	cd ./dist && tar -czf $(NAME)-linux-$(VERSION).tgz $(NAME); rm -f $(NAME)

	GOOS=windows $(BUILD_CMD) -o ./dist/$(NAME).exe .
	cd ./dist && zip $(NAME)-windows-$(VERSION).zip $(NAME).exe; rm -f $(NAME).exe
