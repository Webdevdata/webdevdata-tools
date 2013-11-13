CMD_SRCS = $(wildcard */main.go)
CMDS = $(CMD_SRCS:/main.go=)

BLD_DIR = build

print:
	echo ${CMDS}

all: $(CMDS)

$(CMDS): %: $(BLD_DIR)/%

$(BLD_DIR)/%:
	mkdir -p $(dir $@)
	go build -o $(abspath $@) ./$*

clean:
	rm -rf ${BLD_DIR}

test:
	go test ./...

.PHONY: clean all test

