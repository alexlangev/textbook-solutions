# Exercise 2

```Makefile
.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
		go fmt ./...

vet: fmt
		go vet ./...

build: vet
		go build

clean: 
		go clean
```

## Thoughts

Kinda sick. Why `.Phony`?? Wouldn't these files always be the same though?
