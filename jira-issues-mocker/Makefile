# This how we want to name the binary output
BINARY=beryllium

BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Build=${BUILD}"

.PHONY: default
default: clean install_deps build ;

# Builds the project
build:
	go build ${LDFLAGS} -o bin/${BINARY} cmd/${BINARY}/main.go

# Installs our project dependencies
install_deps::
	go get -u github.com/icrowley/fake

# Cleans our project: deletes binaries
clean:
	if [ -f bin/${BINARY} ] ; then rm bin/${BINARY} ; fi

