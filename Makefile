GOCC := env GOPATH=$(CURDIR) go

all: *.go
	$(GOCC) build -o urlshortener

test:
	$(GOCC) test -v ./...

start: server.PID

server.PID:
	{ ./urlshortener & echo $$! > $@; }

demo: all start
	chmod +x bin/demo.sh
	bin/demo.sh

stop: server.PID
	kill `cat $<` && rm -f $<

.PHONY: test demo start stop
