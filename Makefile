GO ?= go

.PHONY: dev
dev: test lint

.PHONY: test
test:
	set -e; for dir in `find . -type f -name "go.mod"  | sed -r 's@/[^/]+$$@@' | sort | uniq`; do \
	  (set -xe; cd $$dir; $(GO) test -v -cover -race ./...); \
	done

.PHONY: lint
lint:
	golangci-lint run --verbose ./...
