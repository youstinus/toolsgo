####################
# ToolsGo Makefile #
####################

JOBDATE		?= $(shell date -u +%Y-%m-%dT%H%M%SZ)
GIT_REVISION	?= $(shell git rev-parse --short HEAD)
VERSION		?= $(shell git describe --tags --abbrev=0)

LDFLAGS		+= -s -w
LDFLAGS		+= -X github.com/youstinus/toolsgo/pkg/config.appVersion=$(VERSION)
LDFLAGS		+= -X github.com/youstinus/toolsgo/pkg/config.commit=$(GIT_REVISION)
LDFLAGS		+= -X github.com/youstinus/toolsgo/pkg/config.buildTime=$(JOBDATE)

LDFLAGS_LINUX		+= -linkmode external -extldflags -static

BUILDPATH=$(CURDIR)

# Install on windows
iw:
	GOOS=windows CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" github.com/youstinus/toolsgo/cmd/toolsgo

# Install on linux
i:
	GOOS=linux CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" github.com/youstinus/toolsgo/cmd/toolsgo
