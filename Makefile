PROGRAM		= gitpr
DIST		= /usr/local/bin

all:	$(PROGRAM)

$(PROGRAM): gitpr.go
	go build gitpr.go

install:	$(PROGRAM)
	install -s $(PROGRAM) $(DIST)

clean:; rm gitpr
