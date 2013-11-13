CMD_SRCS = $(wildcard */main.go)
ifeq (${GOOS},windows)
	CMDS = $(CMD_SRCS:/main.go=.exe)
else
	CMDS = $(CMD_SRCS:/main.go=)
endif
OSS = darwin linux windows
ARCHS = 386 amd64

BLD_DIR = build
RELEASE_DIR = release

RELEASES = $(foreach arch,$(ARCHS),$(addsuffix -$(arch).tgz,$(CMDS)))

default: clean all

# Release tasks
release: $(OSS)

$(OSS): %: $(addprefix %/,$(ARCHS))

$(foreach os,$(OSS),$(addprefix $(os)/,$(ARCHS))):
	mkdir -p $(RELEASE_DIR)
	$(MAKE) clean
	GOOS=$(@D) GOARCH=$(@F) $(MAKE) all
	mv build webdevdata-tools
	tar czvf $(RELEASE_DIR)/webdevdata-tools-$(@D)-$(@F).tgz webdevdata-tools
	rm -rf webdevdata-tools

$(RELEASE_DIR)/%: $(ARCHS)

# Build tasks
all: $(CMDS)

$(CMDS): %: $(BLD_DIR)/%

$(BLD_DIR)/%:
	mkdir -p $(dir $@)
	go build -o $(abspath $@) ./$(*:.exe=)

clean:
	rm -rf ${BLD_DIR}

test:
	go test ./...

.PHONY: clean all test

