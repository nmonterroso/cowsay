clean: cleanCows
	rm -rf build

cleancows:
	rm -rf cows/cows.go

$(GOPATH)/bin/go-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

cows/cows.go: $(GOPATH)/bin/go-bindata
	go-bindata -o cows/cows.go -pkg cows -prefix cows -ignore cows.go cows/...

cows: cows/cows.go

build: cows/cows.go
	go build -o build/cowsay github.com/nmonterroso/cowsay/cmd

all: cows build