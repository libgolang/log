GHK ?= 'github-key'
TRAVIS_BUILD_NUMBER ?= 1
TRAVIS_COMMIT ?= $(shell git log --format=\%H -1)

SLUG=libgolang/log


build:
	go build

check: $(GOPATH)/bin/deadcode $(GOPATH)/bin/errcheck $(GOPATH)/bin/golint
	# Run Code Tests
	go vet
	errcheck
	golint
	deadcode

unit-test:
	go test

deploy:
	$(eval VERSION = $(shell cat VERSION).$(TRAVIS_BUILD_NUMBER)+$(shell git log --format=%h -1))
	@curl -H "Authorization: Bearer $(GHK)" \
		-d '{ "tag_name": "$(VERSION)", "target_commitish": "$(TRAVIS_COMMIT)", "name": "$(VERSION)", "body": "Automatic Release of $(VERSION)", "draft": false, "prerelease": false }' \
		"https://api.github.com/repos/$(SLUG)/releases"

$(GOPATH)/bin/deadcode:
	go get -u github.com/tsenart/deadcode
$(GOPATH)/bin/errcheck:
	go get -u github.com/kisielk/errcheck
$(GOPATH)/bin/golint:
	go get -u golang.org/x/lint/golint
