# generic build target when the OS doesn't need to be specified
build: clean
	go build -o bin/generic

# Linux-specific build target
build-linux: clean-linux
	go build -o bin/linux/

# OSX-specific build target
build-osx: clean-osx
	go build -o bin/osx/

# Windows-specific build target
build-win: clean-win
	go build -o bin/windows/

# generic clean target when the OS doesn't need to be specified
clean:
	rm -rf bin/generic
	mkdir -p bin/generic

# Linux-specific clean target
clean-linux:
	rm -rf bin/linux
	mkdir -p bin/linux

# OSX-specific clean target
clean-osx:
	rm -rf bin/osx
	mkdir -p bin/osx

# Windows-specific clean target
clean-win:
	rm -rf bin/windows
	mkdir -p bin/windows

# Run all tests
test:
	go test ./...

# Run all tests and ignore any cached results
test-fresh:
	go test -count=1 ./...

# Run all tests with increased verbosity
test-verbose:
	go test -v ./...

# Run all tests with increased verbosity and ignore any cached results
test-verbose-fresh:
	go test -v -count=1 ./...

.PHONY: build build-linux build-osx build-win clean clean-linux clean-osx clean-win test test-fresh test-verbose test-verbose-fresh
