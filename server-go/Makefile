check: test vet gofmt misspell unconvert staticcheck ineffassign unparam
pedantic: check errcheck

run:
	./scripts/runServer.sh

build:
	./scripts/build.sh

buildLatest:
	./scripts/buildLatest.sh

test:
	./scripts/test.sh

benchmark:
	./scripts/benchmark.sh

vet: | test
	./scripts/vet.sh

staticcheck:
	./scripts/staticcheck.sh

misspell:
	./scripts/missspell.sh

unconvert:
	./scripts/unconvert.sh

ineffassign:
	./scripts/ineffassign.sh

unparam:
	./scripts/unparam.sh

errcheck:
	./scripts/errcheck.sh

gofmt:
	./scripts/gofumpt.sh

gofumpt:
	./scripts/gofumpt.sh