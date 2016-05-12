all: fuzz-build

gencorpus:
	go run gencorpus.go

fuzz-build:
	go-fuzz-build github.com/larixsource/suntech/st600

fuzz:
	go-fuzz -bin=./st600-fuzz.zip -workdir=.

