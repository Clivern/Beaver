GO           ?= go
GOFMT        ?= $(GO)fmt
pkgs          = ./...


## install_revive: Install revive for linting.
install_revive:
	@echo ">> Install revive"
	$(GO) get github.com/mgechev/revive


## style: Check code style.
style:
	@echo ">> checking code style"
	@fmtRes=$$($(GOFMT) -d $$(find . -path ./vendor -prune -o -name '*.go' -print)); \
	if [ -n "$${fmtRes}" ]; then \
		echo "gofmt checking failed!"; echo "$${fmtRes}"; echo; \
		echo "Please ensure you are using $$($(GO) version) for formatting code."; \
		exit 1; \
	fi


## check_license: Check if license header on all files.
check_license:
	@echo ">> checking license header"
	@licRes=$$(for file in $$(find . -type f -iname '*.go' ! -path './vendor/*') ; do \
               awk 'NR<=3' $$file | grep -Eq "(Copyright|generated|GENERATED)" || echo $$file; \
       done); \
       if [ -n "$${licRes}" ]; then \
               echo "license header checking failed:"; echo "$${licRes}"; \
               exit 1; \
       fi


## test_short: Run test cases with short flag.
test_short:
	@echo ">> running short tests"
	$(GO) test -short $(pkgs)


## test: Run test cases.
test:
	@echo ">> running all tests"
	$(GO) test -race -cover $(pkgs)


## lint: Lint the code.
lint:
	@echo ">> Lint all files"
	revive -config config.toml -exclude vendor/... -formatter friendly ./...


## format: Format the code.
format:
	@echo ">> formatting code"
	$(GO) fmt $(pkgs)


## vet: Examines source code and reports suspicious constructs.
vet:
	@echo ">> vetting code"
	$(GO) vet $(pkgs)


## dev_run: Run the application main file.
dev_run:
	$(GO) run beaver.go


## prod_run: Build and run the application.
prod_run: build
	./beaver


## coverage: Create HTML coverage report
coverage:
	rm -f coverage.html cover.out
	$(GO) test -coverprofile=cover.out $(pkgs)
	go tool cover -html=cover.out -o coverage.html


## build: Build the application.
build:
	rm -f beaver
	$(GO) build -o beaver beaver.go


## ci: Run all CI tests.
ci: style check_license test vet lint
	@echo "\n==> All quality checks passed"


help: Makefile
	@echo
	@echo " Choose a command run in Beaver:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: help