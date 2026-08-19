
OS_ID = generic
MACHINE = generic

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Linux)
 	OS_ID = Linux_$(UNAME_M)
endif
ifeq ($(UNAME_S),Darwin)
	OS_ID = Darwin_$(UNAME_M)
endif

CODE_NAME = telnet
SOURCES = $(CODE_NAME).go \
	mod/vars/vars.go

BUILT_SOURCES = $(SOURCES)

all:	$(CODE_NAME)_$(OS_ID)\
		install

$(CODE_NAME)_$(OS_ID): $(BUILT_SOURCES)
	@echo "build the $(CODE_NAME)_$(OS_ID) binary..."
	@go build -o bin/$(CODE_NAME)_$(OS_ID) $(CODE_NAME).go
	@strip bin/$(CODE_NAME)_$(OS_ID)

install:
	@echo "Installing the new $(CODE_NAME) binary..."
	@sudo cp bin/$(CODE_NAME)_$(OS_ID) /usr/local/sbin/$(CODE_NAME)
	@sudo chmod 0755 /usr/local/sbin/$(CODE_NAME)
	@sudo chown 0:0 /usr/local/sbin/$(CODE_NAME)

clean:
	@rm -f bin/*$(OS_ID)*
