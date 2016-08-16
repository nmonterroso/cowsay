clean: cleantemplates
	rm -rf build

cleantemplates:
	rm -rf templates/templates.go

$(GOPATH)/bin/go-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

templates/templates.go: $(GOPATH)/bin/go-bindata
	go-bindata -o templates/templates.go -pkg templates -prefix templates -ignore templates.go templates/...

templates: templates/templates.go

build: templates
	go build -o build/cowsay github.com/nmonterroso/cowsay/cmd

all: templates build