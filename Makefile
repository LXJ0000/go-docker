CMD=go
BIN_PATH=bin
APP_NAME=mydocker

all: clean build install

build:
	$(CMD) build -o $(BIN_PATH)/$(APP_NAME) .

install:
	cp $(BIN_PATH)/$(APP_NAME) /usr/bin/$(APP_NAME)
	cp $(BIN_PATH)/$(APP_NAME) /usr/local/bin/$(APP_NAME)

uninstall:
	rm -rf /usr/bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	rm -rf $(BIN_PATH)/$(APP_NAME)

clean: uninstall