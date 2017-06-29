GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
BUILD_IMAGE := vcoin$(if $(GIT_BRANCH_CLEAN),:$(GIT_BRANCH_CLEAN))

dbuild: vcoinimage
	docker run --rm -it -v $(CURDIR):/vcoin $(BUILD_IMAGE) go run main.go

vcoinimage:
	docker build -t $(BUILD_IMAGE) .
	
