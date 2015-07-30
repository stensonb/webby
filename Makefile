GOROOT=/home/bryan/go
GOBIN=$(GOROOT)/bin
GOCMD=$(GOBIN)/go

SOURCES=webby.go
EXECUTABLE=webby

all: $(SOURCES) $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
        GOROOT=$(GOROOT) CGO_ENABLED=0 GOOS=linux $(GOCMD) build -a -installsuffix cgo $(SOURCES)

clean:
	rm -rf $(EXECUTABLE)
