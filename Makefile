PROGRAM_NAME = pomodoro-daemon
INSTALL_DIR = /bin

build:
	cd code && go build -o $(PROGRAM_NAME) main.go

install: build
	cd code && sudo mv $(PROGRAM_NAME) $(INSTALL_DIR)

clean:
	cd code && go clean
	rm -f $(PROGRAM_NAME)

uninstall:
	sudo rm -f $(INSTALL_DIR)/$(PROGRAM_NAME)

.PHONY: build install clean uninstall

